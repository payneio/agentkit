package minds

import (
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/queues"
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
	percepts queues.PerceptQueue,
	actions queues.ActionQueue,
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
