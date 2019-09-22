package sensors

import (
	"agentkit/pkg/agentkit"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type WebAPI struct {
	URL         string
	Method      string
	ContentType string
	Rate        float64
	Out         agentkit.PerceptQueue
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
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}

			percept := &agentkit.Percept{
				Label: `webpage`,
				Data:  string(body),
			}

			sensor.Out.Enqueue(percept)

		}

	}(sensor)

	fmt.Println("WebAPI sensor started.")

}
