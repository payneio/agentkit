package minds

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/queues"
	"fmt"
)

type LoopbackMind struct {
	Percepts queues.PerceptQueue
	Actions  queues.ActionQueue
}

func (m *LoopbackMind) Start() {

	fmt.Println(`Loopback mind is waking.`)

	// agent cycle
	go func(m *LoopbackMind) {
		for {
			if m.Percepts.Peek() != nil {
				percept := m.Percepts.Dequeue()
				action := &datatypes.Action{
					Label: `echo.` + percept.Label,
					Data:  percept.Data,
				}
				m.Actions.Enqueue(action)
			}
		}
	}(m)

}
