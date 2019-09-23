package actuators

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/queues"
	"fmt"
)

// Actuator is anything that can take actions.
type Actuator interface {
	GetLabel() string
	Actuate(*datatypes.Action)
}

type ActuatorConfig struct {
	Type  string
	Label string
}

func New(config *ActuatorConfig, actions queues.ActionQueue) Actuator {

	switch config.Type {
	case `stdout`:
		return &StdOut{
			Label: config.Label,
			In:    actions,
		}
	}

	fmt.Println(`Unknown actuator type.`)
	return nil
}
