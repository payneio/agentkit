package actuators

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
	"os/exec"
)

/*
This actuator allows an agent to speak.
*/

type Speak struct {
	Label string
}

func (a *Speak) GetLabel() string {
	return a.Label
}

func (a *Speak) Actuate(action *datatypes.Action) {

	fmt.Printf("Speaking: %v\n", action.Data)

	phrase, ok := action.Data.(string)
	if !ok {
		return
	}

	// TODO: Need to configure multiple speach systems and check for their
	// existence.
	cmd := exec.Command("espeak", phrase)
	_ = cmd.Run()

	fmt.Println(action)
}
