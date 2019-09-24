package coordinator

import (
	"agentkit/pkg/agentkit/util"
	"fmt"
	"io/ioutil"
	"strconv"

	"cuelang.org/go/cue"
	"github.com/gin-gonic/gin"
)

type Agent struct {
	Name    string
	Address string
}

type Coordinator struct {
	Port   int
	Webd   *gin.Engine
	Agents map[string]*Agent
}

func (c *Coordinator) Spin() {

	// Use the configured port, or find a free one
	if c.Port == 0 {
		c.Port = util.FindFreeTCPPort()
	}
	portStr := strconv.Itoa(c.Port)
	fmt.Println(`Agent Coordinator available on port ` + portStr)
	c.Webd.Run(`:` + portStr)

}

func New(config *cue.Instance) *Coordinator {

	r := gin.Default()

	// Turn off GIN logging
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	port, _ := config.Lookup(`_port`).Int64()
	c := &Coordinator{
		Port:   int(port),
		Webd:   r,
		Agents: make(map[string]*Agent),
	}

	// Set Routes
	r.GET("/", c.Root)
	r.GET("/health", c.Health)
	r.GET("/agents", c.ListAgents)
	r.POST("/agents", c.UpsertAgent)

	return c
}
