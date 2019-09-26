package minds

import (
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
	Beliefs  Beliefs
	Rules    []CARule
}

func (m *CAMind) GetBeliefs() Beliefs {
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
			fmt.Println(`Received percept: ` + percept.Label)
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

	fmt.Println(`Evaluating rule expression: ` + expression)

	// Prepare the environment for condition evaluation.
	// This is done on every condition to allow for simple cascading rules.
	env := map[string]interface{}{
		`beliefs`: m.Beliefs.MSI(),
	}
	fmt.Printf("%#v", env)
	// Evaluate
	out, err := expr.Eval(expression, expr.Env(env))
	if err != nil {
		fmt.Printf("Could not evaluate condition expression. err = %v\n", err)
	}

	// Make sure result is a boolean
	yesno, ok := out.(bool)
	if !ok {
		fmt.Printf("Condition did not evaluate to boolean. expression = %s\n", expression)
	}

	return yesno
}

func (m *CAMind) EvalAction(expression string) {
	if expression == "" {
		return
	}

	// TODO: Create some sort of action language

	fmt.Printf("Action needs to be evaluated: %s\n", expression)

	// Take an action
	// action := &datatypes.Action{
	// 	Label: `echo.` + percept.Label,
	// 	Data:  percept.Data,
	// 	TS:    time.Now(),
	// }
	// m.Actions <- action
}
