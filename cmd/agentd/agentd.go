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
			URL:  `https://api.openweathermap.org/data/2.5/weather?zip=98177,us&units=imperial&APPID=11c411febfa2057a80a18d89ff570383`,
			Rate: 0.1,
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
