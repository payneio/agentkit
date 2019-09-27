package beliefs

import (
	"agentkit/pkg/agentkit/datatypes"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

type BeliefsConfig struct {
	Persistence string
}

type Beliefs interface {
	Perceive(*datatypes.Percept)
	Get(string) *datatypes.Belief
	Set(string, interface{})
	MSI() map[string]interface{}
}

type BasicBeliefs struct {
	config *BeliefsConfig

	// TODO: This is the simplest possible belief repository. The next
	// version should be a tree or a graph. This is not thread-safe, which
	// should be ok for now as we only have one mind.
	facts map[string]*datatypes.Belief
}

// Perceive allows the beliefs to be modified based on the perception. Returns
// whether or not the existing belief has been modified.
func (b *BasicBeliefs) Perceive(p *datatypes.Percept) {
	b.Set(p.Label, p.Data)
}

func (b *BasicBeliefs) Get(key string) *datatypes.Belief {
	belief, ok := b.facts[key]
	if ok {
		return belief
	}
	return nil
}

func (b *BasicBeliefs) Set(key string, val interface{}) {

	now := time.Now()

	belief := &datatypes.Belief{
		ID:        key,
		Data:      val,
		UpdatedAt: now,
		ChangedAt: now,
	}

	// If we already believe this fact, and it hasn't changed, then
	// use the previous changedAt timestamp
	if prev, ok := b.facts[key]; ok {
		if prev.Data == val {
			belief.ChangedAt = prev.ChangedAt
		}

	}

	b.facts[key] = belief

}

func (b *BasicBeliefs) MSI() map[string]interface{} {
	var msi map[string]interface{}
	s, _ := json.Marshal(b.facts)
	json.Unmarshal(s, &msi)
	return msi
}

func NewBasicBeliefs(config *BeliefsConfig) Beliefs {

	log.Info(`I am forming beliefs.`)

	return &BasicBeliefs{
		config: config,
		facts:  make(map[string]*datatypes.Belief),
	}

	return nil
}
