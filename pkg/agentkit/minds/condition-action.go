package minds

import (
	"agentkit/pkg/agentkit/belief"
	"agentkit/pkg/agentkit/datatypes"
	"fmt"

	"github.com/antonmedv/expr"
)

type CARule struct {
	If   string
	Then string
	Else string
}

type CAMind struct {
	Percepts chan *datatypes.Percept
	Actions  chan *datatypes.Action
	Beliefs  *belief.Beliefs
	Rules    []CARule
}

func (m *CAMind) GetBeliefs() *belief.Beliefs {
	return m.Beliefs
}

func (m *CAMind) Start() {

	fmt.Println(`Condition-Action mind is waking.`)

	// agent cycle
	go func(m *CAMind) {

		for {
			percept := <-m.Percepts

			// Form a belief about this percept
			_ = m.Beliefs.Perceive(percept)

			// Eval actions based on whether condition is met or not
			for _, rule := range m.Rules {
				if m.EvalCondition(rule.If) {
					m.EvalAction(rule.Then)
				} else {
					m.EvalAction(rule.Else)
				}
			}

		}

	}(m)

}

func (m *CAMind) EvalCondition(expression string) bool {

	// TODO: Put all stuff needed to evaluate an expression in this (or a struct)
	var env map[string]interface{}

	// Evaluate
	out, err := expr.Eval(expression, env)
	if err != nil {
		fmt.Printf(`Could not evaluate condition expression. err = %v\n`, err)
	}

	// Make sure result is a boolean
	yesno, ok := out.(bool)
	if !ok {
		fmt.Printf(`Condition did not evaluate to boolean. expression = %s\n`, expression)
	}

	return yesno
}

func (m *CAMind) EvalAction(expression string) {
	if expression == "" {
		return
	}

	// TODO: Create some sort of action language

	// Take an action
	// action := &datatypes.Action{
	// 	Label: `echo.` + percept.Label,
	// 	Data:  percept.Data,
	// 	TS:    time.Now(),
	// }
	// m.Actions <- action
}
