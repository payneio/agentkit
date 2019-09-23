package belief

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
)

type Config struct {
	Persistence string
}

type Beliefs struct {
	config *Config

	// TODO: This is the simplest possible belief repository. The next
	// version should be a tree or a graph. This is not thread-safe, which
	// should be ok for now as we only have one mind.
	facts map[string]interface{}
}

func (b *Beliefs) Perceive(p *datatypes.Percept) {
	b.facts[p.Label] = p.Data
}

func (b *Beliefs) MSI() map[string]interface{} {
	return b.facts
}

func New(config *Config) *Beliefs {

	fmt.Println(`I am forming beliefs.`)

	return &Beliefs{
		config: config,
		facts:  make(map[string]interface{}),
	}

	return nil
}
