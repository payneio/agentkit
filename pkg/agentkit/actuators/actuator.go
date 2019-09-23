package actuators

import "agentkit/pkg/agentkit/datatypes"

// Actuator is anything that can take actions.
type Actuator interface {
	GetLabel() string
	Actuate(*datatypes.Action)
}

type ActuatorConfig struct {
	Label string
}
