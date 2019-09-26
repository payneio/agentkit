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
        program: "espeak"
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
            if: "beliefs['weather.temp'] > 76"
            then: "setBelief('outside.hot', true)"
            else: "setBelief('outside.hot, false)"
        },
        {
            if: "beliefs['weather.temp'] < 60"
            then: "setBelief('outside.cold', true)"
            else: "setBelief('outside.cold, false)"
        },
        {
            if: "not(beliefs['outside.cold'] or beliefs['outside.hot'])"
            then: "setBelief('outside.niceTemp', true)"
            else: "setBelief('outside.niceTemp, false)"
        },
        {
            if: "beliefs['weather.windSpeed'] > 15.0"
            then: "setBelief('outside.windy', true)"
            else: "setBelief('outside.windy', false)"
        },
        {
            if: "beliefs['weather.cloudCoverage'] > 80.0"
            then: "setBelief('outside.overcast', true)"
            else: "setBelief('outside.overcast', false)"
        },
        {
            if: "beliefs['outside.hot'] and beliefs['weather.humidity'] > 80 and beliefs['weather.windsSpeed'] < 5.0"
            then: "setBelief('outside.muggy', true)"
            else: "setBelief('outside.muggy', false)"
        },
        {
            if: "beliefs['outside.niceTemp'] and not (beliefs['outside.muggy'] or beliefs['outside.windy'] or beliefs['outside.overcast'])"
            then: "setBelief('outside.comfortable', true)"
            else: "setBelief('outside.comfortable', false)"
        },
        {
            if: "beliefs['outside.comfortable'] == true"
            then: "action('speak', 'It's nice out. You should go outside.')"
            //then: "SMS(textmsg, beliefs.person.paul.phone, 'Go Outside!')"
            else: "action('speak', 'Code away.')"
        }
    ],
    beliefs: [
        "person.paul.phone = 206-790-6707"
    ]
}
