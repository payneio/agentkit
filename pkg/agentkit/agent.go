package agentkit

import "time"

// Sensor is anything that can receive or generate data for the agent.
type Sensor interface {
	Start()
}

// Actuator is anything that can take actions.
type Actuator interface {
	Actuate(Act)
}

// Percept is a datatype required for the percept queue
type Percept struct {
	Label string
	Data  string
	TS    time.Time
}

// Act is an action datatype
type Act struct {
	Label string
	Data  string
	TS    time.Time
}

// Agent is an agent
type Agent struct {
	Sensors   []Sensor
	Actuators []Actuator
}
