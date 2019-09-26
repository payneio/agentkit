package ca

import (
	"testing"

	"github.com/antonmedv/expr"
)

func TestExpr(t *testing.T) {
	expression := "beliefs['temp'] > 60 and beliefs['temp'] < 76"
	env := map[string]interface{}{
		`beliefs`: map[string]interface{}{
			`temp`: 58.6,
		},
	}
	out, err := expr.Eval(expression, env)
	if err != nil {
		t.Errorf("Expected no errors. Got: %v", err)
	}
	if out != false {
		t.Errorf("Expecte `false`, got: %v", out)
	}
}
