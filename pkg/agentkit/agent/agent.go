package agent

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/actuators"
	kactuators "agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/central"
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/minds"
	ksensors "agentkit/pkg/agentkit/sensors"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cuelang.org/go/cue"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// Agent is an agent
type Agent struct {
	webd           *gin.Engine
	Name           string
	Port           int
	PublicAddress  string
	Central        datatypes.Central
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
	agent.MaintainCentralConnection()

	go func() {
		agent.webd.Run(fmt.Sprintf(`:%d`, agent.Port))
	}()
}

func (agent *Agent) Spin() {
	agent.Start()
	select {}
}

func (agent *Agent) MaintainCentralConnection() {
	go func() {
		for {
			agent.NotifyCentral()
			time.Sleep(centralTTL)
		}
	}()
}

func (agent *Agent) NotifyCentral() {

	if agent.Central.Address == "" {
		return
	}

	// Prepare data to POST to Central
	agentData := datatypes.Agent{
		Name:    agent.Name,
		Address: agent.PublicAddress,
		Central: datatypes.Central{
			Name:        agent.Central.Name,
			Address:     agent.Central.Address,
			LastCheckin: agent.Central.LastCheckin,
		},
	}
	agentJSON, _ := json.Marshal(agentData)

	// Post to Central
	url := fmt.Sprintf(`http://%s/agents`, agent.Central.Address)
	result, err := http.Post(url, `application/json`, bytes.NewBuffer(agentJSON))
	if err != nil {
		agent.Central.Status = `lost`
		log.WithFields(log.Fields{`err`: err}).Error(`Failed notifying Central.`)
		return
	}

	// Central returns the populated Agent datatype. We want to update our
	// local info on Central using it.
	data, _ := ioutil.ReadAll(result.Body)
	_ = json.Unmarshal(data, &agentData)
	agent.Central.Name = agentData.Central.Name
	agent.Central.Status = `healthy`
	agent.Central.LastCheckin = agentData.Central.LastCheckin

	log.Info(`Notified Central.`)
}

func New(config *cue.Instance) (*Agent, error) {

	// Agent data
	var agentData *Agent
	err := config.Lookup(`_agent`).Decode(&agentData)
	if err != nil {
		return nil, err
	}
	if agentData.PublicAddress == "" {
		agentData.PublicAddress = fmt.Sprintf(`localhost:%d`, agentData.Port)
	}

	// Channels
	percepts := make(chan *datatypes.Percept)
	actions := make(chan *datatypes.Action)

	// Sensors
	var sensorConfigs []*ksensors.Config
	err = config.Lookup(`sensors`).Decode(&sensorConfigs)
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
		if actuator != nil {
			actuators = append(actuators, actuator)
		}
	}

	// ActionDispatch
	actionDispatch := agentkit.NewActionDispatch(actions)
	actionDispatch.RegisterAll(actuators)

	// Mind
	var mindConfig *minds.Config
	err = config.Lookup(`mind`).Decode(&mindConfig)
	if err != nil {
		return nil, err
	}
	mind := minds.New(mindConfig, percepts, actions)

	// Central connection
	if agentData.Central.Address == "" {
		// If no address, at least check localhost, known port.
		// TODO: Optionally turn this localhost checking off. Nice to
		// have it on unless resource restraints require it to be off.
		agentData.Central.Address = fmt.Sprintf(`localhost:%d`, central.DefaultPort)
	}

	agent := &Agent{
		Name:           agentData.Name,
		Port:           agentData.Port,
		PublicAddress:  agentData.PublicAddress,
		Sensors:        sensors,
		Actuators:      actuators,
		Mind:           mind,
		ActionDispatch: actionDispatch,
		Central:        agentData.Central,
	}

	// JSON Web API
	r := gin.New()
	r.Use(ginlogrus.Logger(log.New()), gin.Recovery())

	// Turn off GIN logging
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Set Routes
	r.GET("/", agent.WebRoot)
	r.GET("/health", agent.WebHealth)
	r.GET("/mind", agent.WebReadMind)

	// Put web server on agent
	agent.webd = r

	return agent, nil
}
