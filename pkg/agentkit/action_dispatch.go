package agentkit

import (
	"agentkit/pkg/agentkit/queues"
	"fmt"
	"strings"
)

type ActionDispatch struct {
	Actions     queues.ActionQueue
	actuatorMap map[string]Actuator
}

func (dispatch *ActionDispatch) Start() {

	go func(dispatch *ActionDispatch) {

		for {

			if dispatch.Actions.Peek() != nil {

				action := dispatch.Actions.Dequeue()
				if action == nil {
					continue
				}

				labelSegs := strings.Split(action.Label, `.`)

				actuatorKey := labelSegs[0]
				actuator := dispatch.actuatorMap[actuatorKey]
				if actuator == nil {
					continue
				}
				actuator.Actuate(action)
			}

		}

	}(dispatch)

}

func (dispatch *ActionDispatch) Register(actuator Actuator) {
	dispatch.actuatorMap[actuator.GetLabel()] = actuator
	fmt.Println(`Registered actuator: ` + actuator.GetLabel())
}

func (dispatch *ActionDispatch) RegisterAll(actuators []Actuator) {
	for _, actuator := range actuators {
		dispatch.Register(actuator)
	}
}

func NewActionDispatch(actions queues.ActionQueue) *ActionDispatch {
	return &ActionDispatch{
		Actions:     actions,
		actuatorMap: make(map[string]Actuator),
	}
}
