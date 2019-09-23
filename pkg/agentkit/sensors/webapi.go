package sensors

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/queues"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/stretchr/objx"
)

type WebAPI struct {
	Config *Config
	Out    queues.PerceptQueue
}

func (sensor *WebAPI) Start() {

	go func(sensor *WebAPI) {

		for {

			sleepDuration := time.Duration((1.0/sensor.Config.Rate)*1000) * time.Millisecond
			time.Sleep(sleepDuration)

			fmt.Println("Reading")

			resp, err := http.Get(sensor.Config.Request.URL)
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Extract JSON
			var m objx.Map
			if sensor.Config.Request.ContentType == `application/json` {
				m = objx.MustFromJSON(string(body))
			}

			if sensor.Config.Measurements != nil {

				for _, measure := range sensor.Config.Measurements {

					// JSONPath Extraction
					if measure.JSONPath != "" {
						v := m.Get(measure.JSONPath).Data()
						percept := &datatypes.Percept{
							Label: sensor.Config.Label + `.` + measure.Value,
							Data:  v,
							TS:    time.Now(),
						}
						sensor.Out.Enqueue(percept)
					}
				}
				continue
			}

			// No specific measurements are requested to be parsed from the
			// body, so go ahead and return the whole body.

			if sensor.Config.Request.ContentType == `application/json` {
				var jsonBody map[string]interface{}
				json.Unmarshal(body, &jsonBody)
				json, err := json.Marshal(jsonBody)
				if err != nil {
					fmt.Println(`Failed marshalling JSON body.`)
				}
				percept := &datatypes.Percept{
					Label: sensor.Config.Label + `.` + `json`,
					Data:  string(json),
					TS:    time.Now(),
				}
				sensor.Out.Enqueue(percept)
			} else {

				percept := &datatypes.Percept{
					Label: sensor.Config.Label + `.` + `webpage`,
					Data:  string(body),
					TS:    time.Now(),
				}
				sensor.Out.Enqueue(percept)
			}

		}

	}(sensor)

	fmt.Println("WebAPI sensor started.")
}

func NewWebAPISensor(config *Config, out queues.PerceptQueue) *WebAPI {
	return &WebAPI{
		Config: config,
		Out:    out,
	}
}
