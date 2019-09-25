package minds

import (
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
	"time"
)

type LoopbackMind struct {
	Percepts chan *datatypes.Percept
	Actions  chan *datatypes.Action
	Beliefs  *belief.Beliefs
}

func (m *LoopbackMind) GetBeliefs() *belief.Beliefs {
	return m.Beliefs
}

func (m *LoopbackMind) Start() {

	fmt.Println(`Loopback mind is waking.`)

	// agent cycle
	go func(m *LoopbackMind) {

		for {
			percept := <-m.Percepts

			// Take an action based on our new perceptions
			action := &datatypes.Action{
				Label: `echo.` + percept.Label,
				Data:  percept.Data,
				TS:    time.Now(),
			}
			m.Actions <- action

			// Form a belief about this percept
			updatedBeliefs := m.Beliefs.Perceive(percept)

			// Optionally, take an action based on changing beliefs.
			if updatedBeliefs {
				//fmt.Println("Updated belief.")
			}
		}

	}(m)

}
