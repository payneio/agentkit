package agentkit

import (
	"fmt"
	"time"
)

// Sensor is anything that can receive or generate data for the agent.
type Sensor interface {
	Start()
}

// Actuator is anything that can take actions.
type Actuator interface {
	GetLabel() string
	Actuate(*Action)
}

// Percept is a datatype required for the percept queue
type Percept struct {
	Label string
	Data  string
	TS    time.Time
}

// Action is an action datatype
type Action struct {
	Label string
	Data  string
	TS    time.Time
}

// Agent is an agent
type Agent struct {
	Sensors        []Sensor
	Actuators      []Actuator
	Mind           Mind
	ActionDispatch *ActionDispatch
}

func (agent *Agent) Start() {

	agent.ActionDispatch.Start()
	agent.Mind.Start()
	fmt.Println(`mind block`)
	for _, sensor := range agent.Sensors {
		sensor.Start()
	}
}

func (agent *Agent) Spin() {
	agent.Start()
	select {}
}
