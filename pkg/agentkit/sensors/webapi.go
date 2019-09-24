package sensors

import (
	"agentkit/pkg/agentkit/datatypes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/stretchr/objx"
)

type WebAPI struct {
	Config *Config
	Out    chan *datatypes.Percept
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

			body, jsonData := s.doHTTP()

			if sensor.Config.Measurements != nil {

				for _, measure := range sensor.Config.Measurements {

					// JSONPath Extraction
					if measure.JSONPath != "" {
						v := jsonData.Get(measure.JSONPath).Data()

						sensor.Out <- &datatypes.Percept{
							Label: sensor.Config.Label + `.` + measure.Value,
							Data:  v,
							TS:    time.Now(),
						}
					}
				}

			} else if sensor.Config.Request.ContentType == `application/json` {

				// No specific measurements are requested to be parsed from the
				// body, so go ahead and return the whole body.

				// As JSON if it is.

				var jsonBody map[string]interface{}
				json.Unmarshal(body, &jsonBody)
				json, err := json.Marshal(jsonBody)
				if err != nil {
					fmt.Println(`Failed marshalling JSON body.`)
				}

				sensor.Out <- &datatypes.Percept{
					Label: sensor.Config.Label + `.` + `json`,
					Data:  string(json),
					TS:    time.Now(),
				}

			} else {

				// Otherwise, just text

				sensor.Out <- &datatypes.Percept{
					Label: sensor.Config.Label + `.` + `webpage`,
					Data:  string(body),
					TS:    time.Now(),
				}
			}

			s.Wait()
		}

	}(s)

	fmt.Println("WebAPI sensor started.")
}

func NewWebAPISensor(config *Config, out chan *datatypes.Percept) *WebAPI {
	return &WebAPI{
		Config: config,
		Out:    out,
	}
}
