package sensors

import (
	"agentkit/pkg/agentkit/queues"
	"fmt"
)

// Sensor is anything that can receive or generate data for the agent.
type Sensor interface {
	Start()
}

type SensorConfig struct {
	Type string
	URL  string
	Rate float64
}

func New(config *SensorConfig, out queues.PerceptQueue) Sensor {

	switch config.Type {
	case `webapi`:
		return &WebAPI{
			URL:  config.URL,
			Rate: config.Rate,
			Out:  out,
		}
	}
	fmt.Println(`Unknown sensor type.`)
	return nil
}
