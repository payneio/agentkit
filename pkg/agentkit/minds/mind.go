package minds

import (
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
	actions queues.ActionQueue) Mind {

	switch config.Type {
	case `loopback`:
		return &LoopbackMind{
			Percepts: percepts,
			Actions:  actions,
		}
	}

	fmt.Println(`Unkown mind type.`)
	return nil
}
