package sensors

import "fmt"

type WebAPI struct {
	URL         string
	Method      string
	ContentType string
	Rate        float64
}

func (sensor *WebAPI) Start() {
	fmt.Println("WebAPI sensor started.")
}
