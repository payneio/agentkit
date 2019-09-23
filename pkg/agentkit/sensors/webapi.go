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

type Values struct {
	Value    string
	JSONPath string
}

type WebAPI struct {
	Label         string
	URL           string
	Method        string
	ContentType   string
	Rate          float64
	ExtractValues []Values
	Out           queues.PerceptQueue
}

func (sensor *WebAPI) Start() {

	go func(sensor *WebAPI) {

		for {

			sleepDuration := time.Duration((1.0/sensor.Rate)*1000) * time.Millisecond
			time.Sleep(sleepDuration)

			fmt.Println("Reading")

			resp, err := http.Get(sensor.URL)
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
			if sensor.ContentType == `application/json` {
				if sensor.ExtractValues != nil {

					for _, value := range sensor.ExtractValues {

						// {"coord":{"lon":-122.37,"lat":47.75},"weather":[{"id":500,"main":"Rain","description":"light
						// rain","icon":"10n"}],"base":"stations","main":{"temp":57.34,"pressure":1017,"humidity":93,"temp_min":55.99,"temp_max":
						// 59},"visibility":16093,"wind":{"speed":5.82,"deg":160},"rain":{"1h":0.7},"clouds":{"all":90},"dt":1569221470,"sys
						// ":{"type":1,"id":5301,"message":0.011,"country":"US","sunrise":1569160561,"sunset":1569204510},"timezone":-25200,
						// "id":0,"name":"Seattle","cod":200} 0001-01-01
						// 00:00:00 +0000 UTC}

						m := objx.MustFromJSON(string(body))
						v := m.Get(value.JSONPath).Data()
						percept := &datatypes.Percept{
							Label: sensor.Label + `.` + value.Value,
							Data:  v,
							TS:    time.Now(),
						}
						sensor.Out.Enqueue(percept)
					}

				} else {
					var jsonBody map[string]interface{}
					json.Unmarshal(body, &jsonBody)
					json, err := json.Marshal(jsonBody)
					if err != nil {
						fmt.Println(`Failed marshalling JSON body.`)
					}
					percept := &datatypes.Percept{
						Label: sensor.Label + `.` + `json`,
						Data:  string(json),
						TS:    time.Now(),
					}
					sensor.Out.Enqueue(percept)
				}

			} else {

				percept := &datatypes.Percept{
					Label: sensor.Label + `.` + `webpage`,
					Data:  string(body),
					TS:    time.Now(),
				}
				sensor.Out.Enqueue(percept)
			}

		}

	}(sensor)

	fmt.Println("WebAPI sensor started.")

}
