perceptQueues: [
    { 
        label: "default",
    }
],
actionQueues: [
    { 
        label: "default",
    }
],
sensors: [
    {
        type: "webapi",
        label: "weather",
        url: "https://api.openweathermap.org/data/2.5/weather?zip=98177,us&units=imperial&APPID=11c411febfa2057a80a18d89ff570383",
        method: "GET",
        contentType: "application/json",
        rate: float64,
        rate: 1.0 / ( 60 * 10 ),
        extractValues: [
            { value: "temp", jsonPath: "main.temp", type: "float64" },
            { value: "pressure", jsonPath: "main.pressure", type: "int" },
            { value: "humidity", jsonPath: "main.humidity", type: "int" },
            { value: "windSpeed", jsonPath: "wind.speed", type: "float64" },
            { value: "windDirection", jsonPath: "wind.deg", type: "int" },
            { value: "cloudCoverage", jsonPath: "clouds.all", type: "int" }
        ]
	    out:  "default",
    },
],
actuators: [
    {
        type: "stdout",
        label: "echo",
		in: "default",
    }
],
mind: {
    type: "loopback"
}
