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

func (s *WebAPI) Wait() {
	sleepDuration := time.Duration((1.0/s.Config.Rate)*1000) * time.Millisecond
	time.Sleep(sleepDuration)
}

func (s *WebAPI) doHTTP() ([]byte, objx.Map) {
	resp, err := http.Get(s.Config.Request.URL)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	// Extract JSON
	var m objx.Map
	if s.Config.Request.ContentType == `application/json` {
		m = objx.MustFromJSON(string(body))
	}

	return body, m
}

func (s *WebAPI) Start() {

	go func(sensor *WebAPI) {

		for {

			fmt.Println("Reading")

			body, jsonData := s.doHTTP()

			if sensor.Config.Measurements != nil {

				for _, measure := range sensor.Config.Measurements {

					// JSONPath Extraction
					if measure.JSONPath != "" {
						v := jsonData.Get(measure.JSONPath).Data()
						percept := &datatypes.Percept{
							Label: sensor.Config.Label + `.` + measure.Value,
							Data:  v,
							TS:    time.Now(),
						}
						sensor.Out.Enqueue(percept)
					}
				}

			} else if sensor.Config.Request.ContentType == `application/json` {

				// No specific measurements are requested to be parsed from the
				// body, so go ahead and return the whole body.

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

			s.Wait()
		}

	}(s)

	fmt.Println("WebAPI sensor started.")
}

func NewWebAPISensor(config *Config, out queues.PerceptQueue) *WebAPI {
	return &WebAPI{
		Config: config,
		Out:    out,
	}
}
