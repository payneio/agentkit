package agent

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/actuators"
	kactuators "agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/minds"
	ksensors "agentkit/pkg/agentkit/sensors"
	"agentkit/pkg/agentkit/util"
	"net/http"
	"strconv"

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

func New(config *cue.Instance) (*Agent, error) {

	// Channels
	percepts := make(chan *datatypes.Percept)
	actions := make(chan *datatypes.Action)

	// Sensors
	var sensorConfigs []*ksensors.Config
	err := config.Lookup(`sensors`).Decode(&sensorConfigs)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}
	mind := minds.New(mindConfig, percepts, actions, beliefs)

	// JSON Web API
	r := render.New()
	gmux := mux.NewRouter()

	gmux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		name, _ := config.Lookup(`_name`).String()
		data := map[string]interface{}{
			`name`: name,
		}
		r.JSON(w, http.StatusOK, data)
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

		// Use the configured port, or find a free one
		port, _ := config.Lookup(`_port`).Int64()
		portStr := strconv.Itoa(int(port))
		if portStr == "" {
			portStr = strconv.Itoa(util.FindFreeTCPPort())
		}

		n.Run(`:` + portStr)
	}()

	return &Agent{
		Sensors:        sensors,
		Actuators:      actuators,
		Mind:           mind,
		ActionDispatch: actionDispatch,
	}, nil
}
