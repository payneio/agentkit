sensorProcessors: [
    {
        type: "sample"
        rate: 1.0
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
    },
],
actuators: [
    {
        type: "stdout"
        label: "echo"
    },
    {
        type: "speak"
        label: "speak"
        config: {
            program: "espeak"
            programConfiguration: {
                voice: "en-us+f3"
            }
        }
    },
    {
        type: "SMS"
        label: "textmsg"
        provider: "T-Mobile"
        number: "2067906707"
    }
],
mind: {
    type: "condition-action",
    rules: [
        {
            if: "belief('weather.temp') > 76"
            then: "setBelief('outside.hot', true)"
            else: "setBelief('outside.hot, false)"
        },
        {
            if: "belief('weather.temp') < 60"
            then: "setBelief('outside.cold', true)"
            else: "setBelief('outside.cold, false)"
        },
        {
            if: "not(belief('outside.cold') or belief('outside.hot'))"
            then: "setBelief('outside.niceTemp', true)"
            else: "setBelief('outside.niceTemp, false)"
        },
        {
            if: "belief('weather.windSpeed') > 15.0"
            then: "setBelief('outside.windy', true)"
            else: "setBelief('outside.windy', false)"
        },
        {
            if: "belief('weather.cloudCoverage') > 80.0"
            then: "setBelief('outside.overcast', true)"
            else: "setBelief('outside.overcast', false)"
        },
        {
            if: "belief('outside.hot') and belief('weather.humidity') > 80 and belief('weather.windsSpeed') < 5.0"
            then: "setBelief('outside.muggy', true)"
            else: "setBelief('outside.muggy', false)"
        },
        {
            if: "belief('outside.niceTemp') and not (belief('outside.muggy') or belief('outside.windy') or belief('outside.overcast'))"
            then: "setBelief('outside.comfortable', true)"
            else: "setBelief('outside.comfortable', false)"
        },
        {
            if: "changed_belief('outside.comfortable') and belief('outside.comfortable') == true"
            then: "action('speak', 'It's nice out. You should go outside.')"
        },
        {
            if: "changed_belief('outside.comfortable') and belief('outside.comfortable') == false"
            then: "action('speak', 'Code away.')"
        }
    ],
    beliefs: [
        "person.paul.phone = 206-790-6707"
    ]
}
