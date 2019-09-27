package loopback

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/minds/beliefs"
	"time"

	log "github.com/sirupsen/logrus"
)

type Mind struct {
	Percepts chan *datatypes.Percept
	Actions  chan *datatypes.Action
	Beliefs  beliefs.Beliefs
}

func (m *Mind) GetBeliefs() beliefs.Beliefs {
	return m.Beliefs
}

func (m *Mind) Start() {

	log.Info(`Loopback mind is waking.`)

	// agent cycle
	go func(m *Mind) {

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
			m.Beliefs.Perceive(percept)

		}

	}(m)

}
