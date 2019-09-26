package beliefs

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
)

type BeliefsConfig struct {
	Persistence string
}

type Beliefs interface {
	Perceive(*datatypes.Percept) bool
	Get(string) interface{}
	Set(string, interface{})
	MSI() map[string]interface{}
}

type BasicBeliefs struct {
	config *BeliefsConfig

	// TODO: This is the simplest possible belief repository. The next
	// version should be a tree or a graph. This is not thread-safe, which
	// should be ok for now as we only have one mind.
	facts map[string]interface{}
}

// Perceive allows the beliefs to be modified based on the perception. Returns
// whether or not the existing belief has been modified.
func (b *BasicBeliefs) Perceive(p *datatypes.Percept) (modified bool) {
	modified = false
	if b.facts[p.Label] != p.Data {
		modified = true
	}
	b.facts[p.Label] = p.Data
	return
}

func (b *BasicBeliefs) Get(key string) interface{} {
	val, ok := b.facts[key]
	if ok {
		return val
	} else {
		return nil
	}
}

func (b *BasicBeliefs) Set(key string, val interface{}) {
	b.facts[key] = val
}

func (b *BasicBeliefs) MSI() map[string]interface{} {
	return b.facts
}

func NewBasicBeliefs(config *BeliefsConfig) Beliefs {

	fmt.Println(`I am forming beliefs.`)

	return &BasicBeliefs{
		config: config,
		facts:  make(map[string]interface{}),
	}

	return nil
}
