package minds

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/minds/beliefs"
	"agentkit/pkg/agentkit/minds/ca"
	"agentkit/pkg/agentkit/minds/loopback"
	"fmt"
)

type Mind interface {
	Start()
	GetBeliefs() beliefs.Beliefs
}

type Config struct {
	Type  string
	Rules []ca.CARule
}

func New(
	config *Config,
	percepts chan *datatypes.Percept,
	actions chan *datatypes.Action) Mind {

	switch config.Type {
	case `condition-action`:
		return &ca.Mind{
			Percepts: percepts,
			Actions:  actions,
			Rules:    config.Rules,
			Beliefs:  beliefs.NewBasicBeliefs(nil),
		}
	case `loopback`:
		return &loopback.Mind{
			Percepts: percepts,
			Actions:  actions,
			Beliefs:  beliefs.NewBasicBeliefs(nil),
		}
	}

	fmt.Println(`Unkown mind type.`)
	return nil
}
