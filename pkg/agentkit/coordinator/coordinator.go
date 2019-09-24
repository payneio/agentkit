package coordinator

import (
	"agentkit/pkg/agentkit/agent"
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/util"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"cuelang.org/go/cue"
	"github.com/gin-gonic/gin"
)

type Agent struct {
	Name    string
	Address string
}

type Coordinator struct {
	Name          string
	PublicAddress string
	Port          int
	Webd          *gin.Engine
	Agents        map[string]*datatypes.Agent
}

func (c *Coordinator) Spin() {

	c.startAgentLivenessCheck()

	// Use the configured port, or find a free one
	if c.Port == 0 {
		c.Port = util.FindFreeTCPPort()
	}
	portStr := strconv.Itoa(c.Port)
	fmt.Println(`Agent Coordinator available on port ` + portStr)
	c.Webd.Run(`:` + portStr)

}

func (c *Coordinator) startAgentLivenessCheck() {

	staleDuration := time.Duration(1 * time.Minute)
	expireDuration := time.Duration(1 * time.Hour)

	go func() {
		now := time.Now()
		for name, agent := range c.Agents {

			if now.After(agent.Coordinator.LastCheckin.Add(staleDuration)) {
				agent.Coordinator.Status = `stale`
			}
			if now.After(agent.Coordinator.LastCheckin.Add(expireDuration)) {
				delete(c.Agents, name)
			}

		}
		time.Sleep(10 * time.Second)
	}()
}

func New(config *cue.Instance) *Coordinator {

	r := gin.Default()

	// Turn off GIN logging
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	port, _ := config.Lookup(`_port`).Int64()
	c := &Coordinator{
		Name:          agent.GenerateName(),
		PublicAddress: `localhost:` + strconv.Itoa(int(port)),
		Port:          int(port),
		Webd:          r,
		Agents:        make(map[string]*datatypes.Agent),
	}

	// Set Routes
	r.GET("/", c.Root)
	r.GET("/health", c.Health)
	r.GET("/agents", c.ListAgents)
	r.POST("/agents", c.UpsertAgent)

	return c
}
