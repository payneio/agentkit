package sensors

import "agentkit/pkg/agentkit/queues"

// Sensor is anything that can receive or generate data for the agent.
type Sensor interface {
	Start()
}

type SensorConfig struct {
	URL  string
	Rate float64
}

func New(config *SensorConfig, out queues.PerceptQueue) Sensor {

	// TODO: type
	return &WebAPI{
		URL:  config.URL,
		Rate: config.Rate,
		Out:  out,
	}
}
