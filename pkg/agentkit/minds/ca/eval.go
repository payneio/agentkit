package ca

import (
	"agentkit/pkg/agentkit/datatypes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"regexp"

	"github.com/antonmedv/expr"
)

func (m *Mind) EvalCondition(expression string) bool {

	// Prepare the environment for condition evaluation.
	// This is done on every condition to allow for simple cascading rules.
	env := map[string]interface{}{
		`beliefs`: m.Beliefs.MSI(),
		`belief`: func(key string) interface{} {
			belief := m.Beliefs.Get(key)
			if belief == nil {
				return nil
			}
			return belief.Data
		},
		`changed_belief`: func(key string) bool {
			belief := m.Beliefs.Get(key)
			if belief == nil {
				return false
			}
			if time.Now().Before(belief.ChangedAt.Add(1 * time.Second)) {
				return true
			}
			return false
		},
	}

	// Evaluate
	out, err := expr.Eval(expression, env)
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

func (m *Mind) EvalAction(expression string) {
	if expression == "" {
		return
	}

	// TODO: Create some sort of action language
	// Simple one for now as an example:

	// Parse Action-string for `setBelief` actions
	re := regexp.MustCompile(`\s*([^(]*)\([\s']*([^']*)[\s']*,\s*(.*)\s*\)`)
	matches := re.FindAllStringSubmatch(expression, -1)
	for _, match := range matches {

		action, label, sval := match[1], match[2], match[3]

		// Basic JSON-ish typing of the value

		// FIXME: Currently have single quotes in config'd rules. But they
		// are not handled like double quotes w/ JSON unmarshal. This just
		// requotes the un-nested JSON value.
		if strings.HasPrefix(sval, `'`) {
			sval = fmt.Sprintf(`"%s"`, strings.Trim(sval, `'`))
		}

		var val interface{}
		json.Unmarshal([]byte(sval), &val)

		switch action {
		case `setBelief`:
			m.Beliefs.Set(label, val)
		case `action`:
			// Take an action
			action := &datatypes.Action{
				Label: label,
				Data:  val,
				TS:    time.Now(),
			}
			m.Actions <- action
		}

	}
}
