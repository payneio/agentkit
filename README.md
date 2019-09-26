# agentkit

## Build

```bash
source source_me
make clean && build
```

## Usage

### Starting your first agent

Run `agent -h` to see the options for running an agent.

Minimally, you need to specify a configuration file. For example:

`agent -c examples/webapi_echo/webapi_echo.cue`

This will start an agent that has a WebAPI sensor that calls out to a
weather API, and a simple (the simplest really) "loopback" mind that doesn't
think about anything, just echos back what it sees, which, in this case
are current weather measurements in Seattle.

You can see what is going on with the agent by visiting it's exposes web API
at the port the agent chose when starting up. You can also specify a particular
port:

`agent -c examples/webapi_echo/webapi_echo.cue -p 9200`

Then visit `http://localhost:9200/` to see what the agent has been named
(also overridable) and other info about the agent.

Visit `http://localhost:9200/mind` to see what is on the agent's mind. In the
example, you will see current weather readings.

### Starting Agent Central

For multi-agent coordination, you can start a "Central":

`agentcentral`

With no parameters, Central starts up on port `9200`. Unless specified otherwise,
agents also look for Central at `http://localhost:9200`, so simply starting up
Central on your local machine will result in any running agents connecting to 
it automatically. You can see which agents are connected:

`http://localhost:9200/agents`

## Agent Design

Pretty simple design:

```
sensor(s) -> mind -> actuators
```

The main point is that everything should be configurable. This enables
running wildly different agents by simply starting them with different
configurations.

## RoadMap

* Makefile needs to detect changed src files
* Get a proper logger
* Add a condition-action mind, possibly using `expr`.
  * Add an email or SMS actuator.
  * Update the config to send an email or SMS when it is nice out.
  * More rich CA action specification
    * a syntax for executing actions
    * belief changes need to understand whether it just changed now, or 
      has been the same in previous cycles... or handle that with a percept
* Add a sensor filter.
* agent2agent communication.
* An agent doesn't necessarily need a coordinator... just another agent and 
  we can use a gossip protocol to find the other data... but maybe not an
  appropriate design.
* Add more minds:
  * subsumption.
  * rules/production engine.
  * computational logic (Swift Prolog).
  * Enable ML models.
  * Contemplate composing minds, e.g. some video stream object recognition
    combined with logic.
* More sensors
  * POST endpoint (for callbacks from 3rd party services, e.g. Slack)
  * BLE proximity
  * Speech input
  * Video streams
* More actuators
  * Arduino/pyyaml/something for a series of servos
  * Speech
  * PLC
* ...
