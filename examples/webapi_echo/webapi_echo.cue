topics: {
    percepts: { },
    actions: { },
}

perceptQueues: [
    { 
        label: "default",
        transport: "inproc",
    }
],
actionQueues: [
    { 
        label: "default",
        transport: "inproc",
    }
],
sensorProcessors: [
    {
        type: "sample"
        rate: 1.0
        plumbing: {
            subscribe: "weather.temp"
            publish: "percepts"
            transport: "inproc"
        }
    }
],
sensors: [
    {
        label: "weather"
        type: "webapi"
        request: {
            url: "https://api.openweathermap.org/data/2.5/weather?zip=98177,us&units=imperial&APPID=11c411febfa2057a80a18d89ff570383"
            method: "GET"
            contentType: "application/json"
        },
        measurements: [
            { value: "temp", jsonPath: "main.temp", type: "gauge", datatype: "float64",
              filters: [{ type: "smooth", "timeWindow": 10, magnitude: 0.90 }]
            },
            { value: "pressure", jsonPath: "main.pressure", type: "gauge", datatype: "int" },
            { value: "humidity", jsonPath: "main.humidity", type: "gauge", datatype: "int" },
            { value: "windSpeed", jsonPath: "wind.speed", type: "gauge", datatype: "float64" },
            { value: "windDirection", jsonPath: "wind.deg", type: "gauge", datatype: "int" },
            { value: "cloudCoverage", jsonPath: "clouds.all", type: "gauge", datatype: "int" },
        ]
        rate: 1.0 / ( 60 * 10 )
        cache expire: 300, // 5 min
        filters: {},
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
    percepts: "percept",
    actions: "action"
}
