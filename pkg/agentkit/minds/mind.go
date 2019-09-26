package minds

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
)

type Mind interface {
	Start()
	GetBeliefs() Beliefs
}

type Config struct {
	Type  string
	Rules []CARule
}

func New(
	config *Config,
	percepts chan *datatypes.Percept,
	actions chan *datatypes.Action) Mind {

	switch config.Type {
	case `condition-action`:
		return &CAMind{
			Percepts: percepts,
			Actions:  actions,
			Rules:    config.Rules,
			Beliefs:  NewBasicBeliefs(nil),
		}
	case `loopback`:
		return &LoopbackMind{
			Percepts: percepts,
			Actions:  actions,
			Beliefs:  NewBasicBeliefs(nil),
		}
	}

	fmt.Println(`Unkown mind type.`)
	return nil
}
