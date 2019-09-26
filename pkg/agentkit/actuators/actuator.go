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

// TODO: Should maybe get rid of this and just use a map[string]interface{}.
// New can check to make sure a type and label are set, then everything
// else can just be what it is.
type ActuatorConfig struct {
	Type   string
	Label  string
	Config map[string]interface{}
}

func New(config *ActuatorConfig, actions chan *datatypes.Action) Actuator {

	switch config.Type {
	case `stdout`:
		return &StdOut{
			Label: config.Label,
			In:    actions,
		}
	case `speak`:
		program, ok := config.Config[`program`]
		// TODO: Can avoid this with CUE validation
		if !ok {
			fmt.Println("Speak actuator must have a program set.")
		}

		var programConfig map[string]interface{}
		if programData, cok := config.Config[`programConfiguration`]; cok {
			var tok bool
			if programConfig, tok = programData.(map[string]interface{}); !tok {
				fmt.Println("Invalid speak program config.")
			}
		}

		if programStr, sok := program.(string); sok {
			return &Speak{
				Label:                config.Label,
				Program:              programStr,
				ProgramConfiguration: programConfig,
			}
		}
		return nil
	}

	fmt.Println(`Unknown actuator type: ` + config.Type)
	return nil
}
