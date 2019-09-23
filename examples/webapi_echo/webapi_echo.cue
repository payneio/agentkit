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
		rate: 0.1,
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
