package minds

import (
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
)

type Mind interface {
	Start()
	GetBeliefs() *belief.Beliefs
}

type Config struct {
	Type  string
	Rules []CARule
}

func New(
	config *Config,
	percepts chan *datatypes.Percept,
	actions chan *datatypes.Action,
	beliefs *belief.Beliefs) Mind {

	switch config.Type {
	case `condition-action`:
		return &CAMind{
			Percepts: percepts,
			Actions:  actions,
			Beliefs:  beliefs,
			Rules:    config.Rules,
		}
	case `loopback`:
		return &LoopbackMind{
			Percepts: percepts,
			Actions:  actions,
			Beliefs:  beliefs,
		}
	}

	fmt.Println(`Unkown mind type.`)
	return nil
}
