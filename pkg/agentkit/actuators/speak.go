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
	Label                string
	Program              string
	ProgramConfiguration map[string]interface{}
}

func (a *Speak) GetLabel() string {
	return a.Label
}

func (a *Speak) Actuate(action *datatypes.Action) {

	phrase, ok := action.Data.(string)
	if !ok {
		return
	}

	// TODO: Need to configure multiple speach systems and check for their
	// existence.
	switch a.Program {
	case `espeak`:
		voice := `default`
		if voiceConfig, ok := a.ProgramConfiguration[`voice`]; ok {
			voice = voiceConfig.(string)
		}
		cmd := exec.Command("espeak", "-v", voice, phrase)
		_ = cmd.Run()
	default:
		// Voice configuration program not recognized... just echo
		fmt.Println(action)
	}

}
