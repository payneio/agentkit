package central

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *Central) Root(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"name":    c.Name,
		"version": "0.0.1",
	})
}

func (c *Central) Health(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func (c *Central) ListAgents(gctx *gin.Context) {

	type agentResponse struct {
		Name             string    `json:"name"`
		Address          string    `json:"address"`
		LastCheckin      time.Time `json:"lastCheckin"`
		ConnectionStatus string    `json:"connectionStatus"`
	}

	agents := []agentResponse{}

	for _, agent := range c.Agents {
		agents = append(agents, agentResponse{
			Name:             agent.Name,
			Address:          agent.Address,
			LastCheckin:      agent.Central.LastCheckin,
			ConnectionStatus: agent.Central.Status,
		})
	}
	// TODO: sort by name, I suppose
	gctx.JSON(200, agents)
}

func (c *Central) UpsertAgent(gctx *gin.Context) {

	// TODO: auth
	// When posting for the first time, central should generate keys
	// and give agent the public key. All subsequent requests should use
	// said public key.

	var agent datatypes.Agent
	if err := gctx.ShouldBindJSON(&agent); err != nil {
		err := fmt.Errorf(`Invalid request body. err = %s`, err)
		code := 401
		gctx.AbortWithStatusJSON(code, err)
		return
	}
	agent.Central = datatypes.Central{
		Name:        c.Name,
		Address:     c.PublicAddress,
		LastCheckin: time.Now(),
		Status:      `healthy`,
	}

	c.Agents[agent.Name] = &agent

	gctx.JSON(200, agent)
}
