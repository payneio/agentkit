package actuators

import (
	"agentkit/pkg/agentkit/datatypes"
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

func New(config *ActuatorConfig, actions chan *datatypes.Action) Actuator {

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
