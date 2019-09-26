package actuators

import (
	"agentkit/pkg/agentkit/datatypes"
	"testing"
)

func TestSpeak(t *testing.T) {
	voice := &Speak{}
	action := datatypes.Action{
		Data: "Hello.",
	}
	voice.Actuate(&action)
}
