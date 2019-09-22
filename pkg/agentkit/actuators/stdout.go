package actuators

import (
	"agentkit/pkg/agentkit"
	"fmt"
)

type StdOut struct {
	Label string
	In    agentkit.ActionQueue
}

func (a *StdOut) GetLabel() string {
	return a.Label
}

func (a *StdOut) Actuate(action *agentkit.Action) {
	fmt.Println(action)
}
