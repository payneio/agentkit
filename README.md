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
