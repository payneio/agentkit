package main

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/sensors"
	"fmt"
)

func main() {
	fmt.Println("MCP initializing.")

	percepts := agentkit.NewInMemoryPerceptQueue()
	actions := agentkit.NewInMemoryActionQueue()

	sensors := []agentkit.Sensor{
		&sensors.WebAPI{
			URL:  "https://google.com",
			Rate: 1.0,
			Out:  percepts,
		},
	}

	actuators := []agentkit.Actuator{
		&actuators.StdOut{
			Label: `echo`,
			In:    actions,
		},
	}

	mind := &agentkit.LoopbackMind{
		Percepts: percepts,
		Actions:  actions,
	}

	actionDispatch := agentkit.NewActionDispatch(actions)
	actionDispatch.RegisterAll(actuators)

	agent := &agentkit.Agent{
		Sensors:        sensors,
		Actuators:      actuators,
		Mind:           mind,
		ActionDispatch: actionDispatch,
	}

	agent.Spin()

	fmt.Println("MCP Ready.")

}
