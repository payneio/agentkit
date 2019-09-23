package minds

import (
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
)

type Mind interface {
	Start()
}

type Config struct {
	Type string
}

func New(
	config *Config,
	percepts chan *datatypes.Percept,
	actions chan *datatypes.Action,
	beliefs *belief.Beliefs) Mind {

	switch config.Type {
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
