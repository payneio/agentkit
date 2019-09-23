package sensors

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
)

// Sensor is anything that can receive or generate data for the agent.
type Sensor interface {
	Start()
}

type ConfigMeasurements struct {
	Value    string
	Type     string
	Datatype string
	JSONPath string
}

type ConfigRequest struct {
	URL         string
	Method      string
	ContentType string
}

type Config struct {
	Type         string
	Request      ConfigRequest
	Rate         float64
	Label        string
	Measurements []ConfigMeasurements
}

func New(config *Config, out chan *datatypes.Percept) Sensor {

	switch config.Type {
	case `webapi`:
		return NewWebAPISensor(config, out)
	}
	fmt.Println(`Unknown sensor type.`)
	return nil
}
