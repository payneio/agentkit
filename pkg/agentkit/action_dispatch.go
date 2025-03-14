package agentkit

import (
	"agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/datatypes"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ActionDispatch struct {
	Actions     chan *datatypes.Action
	actuatorMap map[string]actuators.Actuator
}

func (dispatch *ActionDispatch) Start() {

	go func(dispatch *ActionDispatch) {

		for {
			action := <-dispatch.Actions
			log.WithFields(log.Fields{`action`: action}).Info("Taking action.")

			labelSegs := strings.Split(action.Label, `.`)

			actuatorKey := labelSegs[0]
			actuator := dispatch.actuatorMap[actuatorKey]
			if actuator == nil {
				log.WithFields(log.Fields{`name`: actuatorKey}).Error("No actuator with this name.")
				continue
			}
			actuator.Actuate(action)
		}

	}(dispatch)

}

func (dispatch *ActionDispatch) Register(actuator actuators.Actuator) {
	dispatch.actuatorMap[actuator.GetLabel()] = actuator
	log.Info(`Registered actuator: ` + actuator.GetLabel())
}

func (dispatch *ActionDispatch) RegisterAll(actuators []actuators.Actuator) {
	for _, actuator := range actuators {
		dispatch.Register(actuator)
	}
}

func NewActionDispatch(actions chan *datatypes.Action) *ActionDispatch {
	return &ActionDispatch{
		Actions:     actions,
		actuatorMap: make(map[string]actuators.Actuator),
	}
}
