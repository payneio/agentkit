package agentkit

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/sensors"
	"fmt"
)

// Sensor is anything that can receive or generate data for the agent.
// Actuator is anything that can take actions.
type Actuator interface {
	GetLabel() string
	Actuate(*datatypes.Action)
}

type ActuatorConfig struct {
	Label string
}

// Agent is an agent
type Agent struct {
	Sensors        []sensors.Sensor
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
