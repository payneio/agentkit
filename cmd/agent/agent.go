package main

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/sensors"
	"fmt"
)

func main() {
	fmt.Println("MCP initializing.")

	agent := &agentkit.Agent{
		Sensors: []agentkit.Sensor{
			&sensors.WebAPI{
				URL:  "https://google.com",
				Rate: 0.1,
			},
		},
	}

	for _, sensor := range agent.Sensors {
		sensor.Start()
	}

	fmt.Println("MCP Ready.")
}
