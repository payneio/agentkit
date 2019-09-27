package central

import (
	"agentkit/pkg/agentkit/datatypes"
	"agentkit/pkg/agentkit/util"
	"io/ioutil"
	"strconv"
	"time"

	"cuelang.org/go/cue"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Agent struct {
	Name    string
	Address string
}

type Central struct {
	Name          string
	PublicAddress string
	Port          int
	Webd          *gin.Engine
	Agents        map[string]*datatypes.Agent
}

func (c *Central) Spin() {

	c.startAgentLivenessCheck()

	// Start web server in another thread
	go func() {
		// Use the configured port, or find a free one
		if c.Port == 0 {
			c.Port = util.FindFreeTCPPort()
		}
		portStr := strconv.Itoa(c.Port)
		log.Info(`Central available on port ` + portStr)
		c.Webd.Run(`:` + portStr)
	}()

	// Spin
	select {}

}

func (c *Central) startAgentLivenessCheck() {

	go func() {
		for {
			now := time.Now()
			for name, agent := range c.Agents {
				log.Info(`liveness checking ` + agent.Name)

				if now.After(agent.Central.LastCheckin.Add(staleDuration)) {
					agent.Central.Status = `stale`
				}
				if now.After(agent.Central.LastCheckin.Add(expireDuration)) {
					delete(c.Agents, name)
				}

			}
			time.Sleep(livenessCheckDuration)
		}
	}()
}

func New(config *cue.Instance) *Central {

	r := gin.Default()

	// Turn off GIN logging
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	port, _ := config.Lookup(`_port`).Int64()
	c := &Central{
		Name:          util.GenerateName(),
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
