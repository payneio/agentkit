package actuators

import (
	"agentkit/pkg/agentkit/datatypes"

	log "github.com/sirupsen/logrus"
)

type StdOut struct {
	Label string
	In    chan *datatypes.Action
}

func (a *StdOut) GetLabel() string {
	return a.Label
}

func (a *StdOut) Actuate(action *datatypes.Action) {
	log.Info(action)
}
