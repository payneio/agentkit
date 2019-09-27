package ca

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/minds/beliefs"

	log "github.com/sirupsen/logrus"
)

type CARule struct {
	If   string
	Then string
	Else string
}

type Mind struct {
	Percepts chan *datatypes.Percept
	Actions  chan *datatypes.Action
	Beliefs  beliefs.Beliefs
	Rules    []CARule
}

func (m *Mind) GetBeliefs() beliefs.Beliefs {
	return m.Beliefs
}

func (m *Mind) Start() {

	log.Info(`Condition-Action mind is waking.`)

	// agent cycle
	go func(m *Mind) {

		for {
			percept := <-m.Percepts

			// Form a belief about this percept
			m.Beliefs.Perceive(percept)

			// Eval actions based on whether condition is met or not
			for _, rule := range m.Rules {
				if m.EvalCondition(rule.If) {
					m.EvalAction(rule.Then)
				} else {
					m.EvalAction(rule.Else)
				}
			}

		}

	}(m)

}
