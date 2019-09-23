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
	Type          string
	URL           string
	Rate          float64
	Label         string
	Method        string
	ContentType   string
	ExtractValues []Values
}

func New(config *SensorConfig, out queues.PerceptQueue) Sensor {

	switch config.Type {
	case `webapi`:
		return &WebAPI{
			URL:           config.URL,
			Rate:          config.Rate,
			Label:         config.Label,
			Method:        config.Method,
			ContentType:   config.ContentType,
			ExtractValues: config.ExtractValues,
			Out:           out,
		}
	}
	fmt.Println(`Unknown sensor type.`)
	return nil
}
