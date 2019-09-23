package agentkit

import (
	"agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/minds"
	"agentkit/pkg/agentkit/sensors"
	"fmt"
)

// Agent is an agent
type Agent struct {
	Sensors        []sensors.Sensor
	Actuators      []actuators.Actuator
	Mind           minds.Mind
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
