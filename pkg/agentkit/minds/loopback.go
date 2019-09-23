package minds

import (
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/queues"
	"fmt"
	"time"
)

type LoopbackMind struct {
	Percepts queues.PerceptQueue
	Actions  queues.ActionQueue
	Beliefs  *belief.Beliefs
}

func (m *LoopbackMind) Start() {

	fmt.Println(`Loopback mind is waking.`)

	// agent cycle
	go func(m *LoopbackMind) {
		for {
			if m.Percepts.Peek() != nil {

				// Perceive the world
				percept := m.Percepts.Dequeue()

				// Form a belief about this percept
				m.Beliefs.Perceive(percept)

				// Take an action based on our new perceptions or beliefs
				action := &datatypes.Action{
					Label: `echo.` + percept.Label,
					Data:  percept.Data,
					TS:    time.Now(),
				}
				m.Actions.Enqueue(action)
			}
		}
	}(m)

}
