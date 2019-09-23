package sensors

import "agentkit/pkg/agentkit/queues"

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
