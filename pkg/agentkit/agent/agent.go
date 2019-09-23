package agent

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/actuators"
	kactuators "agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/minds"
	"agentkit/pkg/agentkit/queues"
	ksensors "agentkit/pkg/agentkit/sensors"
	"fmt"
	"net/http"
	"os"

	"cuelang.org/go/cue"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// Agent is an agent
type Agent struct {
	Sensors        []ksensors.Sensor
	Actuators      []actuators.Actuator
	Mind           minds.Mind
	ActionDispatch *agentkit.ActionDispatch
}

func (agent *Agent) Start() {
	agent.ActionDispatch.Start()
	agent.Mind.Start()
	for _, sensor := range agent.Sensors {
		sensor.Start()
	}
}

func (agent *Agent) Spin() {
	agent.Start()
	select {}
}

func New(config *cue.Instance) *Agent {
	// Queues
	percepts := queues.NewInMemoryPerceptQueue()
	actions := queues.NewInMemoryActionQueue()

	// Sensors
	var sensorConfigs []*ksensors.Config
	err := config.Lookup(`sensors`).Decode(&sensorConfigs)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	sensors := []ksensors.Sensor{}
	for _, sensorConfig := range sensorConfigs {
		sensor := ksensors.New(sensorConfig, percepts)
		sensors = append(sensors, sensor)
	}

	// Actuators
	var actuatorConfigs []*actuators.ActuatorConfig
	err = config.Lookup(`actuators`).Decode(&actuatorConfigs)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	actuators := []actuators.Actuator{}
	for _, actuatorConfig := range actuatorConfigs {
		actuator := kactuators.New(actuatorConfig, actions)
		actuators = append(actuators, actuator)
	}

	// ActionDispatch
	actionDispatch := agentkit.NewActionDispatch(actions)
	actionDispatch.RegisterAll(actuators)

	// Beliefs
	beliefs := belief.New(&belief.Config{})

	// Mind
	var mindConfig *minds.Config
	err = config.Lookup(`mind`).Decode(&mindConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	mind := minds.New(mindConfig, percepts, actions, beliefs)

	// JSON Web API
	r := render.New()
	gmux := mux.NewRouter()

	gmux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	gmux.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		var response = map[string]string{"health": "ok"}
		r.JSON(w, http.StatusOK, response)
	})

	gmux.HandleFunc("/version", func(w http.ResponseWriter, req *http.Request) {
		var response = map[string]string{"version": "0.0.1"}
		r.JSON(w, http.StatusOK, response)
	})

	gmux.HandleFunc("/beliefs", func(w http.ResponseWriter, req *http.Request) {
		var response = beliefs.MSI()
		r.JSON(w, http.StatusOK, response)
	})

	go func() {
		n := negroni.Classic()
		n.UseHandler(gmux)
		n.Run(":3000")
	}()

	return &Agent{
		Sensors:        sensors,
		Actuators:      actuators,
		Mind:           mind,
		ActionDispatch: actionDispatch,
	}
}
