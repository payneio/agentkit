package actuators

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/queues"
	"fmt"
)

type StdOut struct {
	Label string
	In    queues.ActionQueue
}

func (a *StdOut) GetLabel() string {
	return a.Label
}

func (a *StdOut) Actuate(action *datatypes.Action) {
	fmt.Println(action)
}
