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

// Perceive allows the beliefs to be modified based on the perception. Returns
// whether or not the existing belief has been modified.
func (b *Beliefs) Perceive(p *datatypes.Percept) (modified bool) {
	modified = false
	if b.facts[p.Label] != p.Data {
		modified = true
	}
	b.facts[p.Label] = p.Data
	return
}

func (b *Beliefs) Get(key string) interface{} {
	val, ok := b.facts[key]
	if ok {
		return val
	} else {
		return nil
	}
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
