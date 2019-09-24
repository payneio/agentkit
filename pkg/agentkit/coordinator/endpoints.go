package coordinator

import (
	"agentkit/pkg/agentkit/datatypes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *Coordinator) Root(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"name":    "agent_coordinator",
		"version": "0.0.1",
	})
}

func (c *Coordinator) Health(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func (c *Coordinator) ListAgents(gctx *gin.Context) {
	agents := []datatypes.Agent{}
	for _, agent := range c.Agents {
		agents = append(agents, *agent)
	}
	// TODO: sort by name, I suppose
	gctx.JSON(200, agents)
}

func (c *Coordinator) UpsertAgent(gctx *gin.Context) {

	// TODO: auth
	// When posting for the first time, coordinator should generate keys
	// and give agent the public key. All subsequent requests should use
	// said public key.

	var agent datatypes.Agent
	if err := gctx.ShouldBindJSON(&agent); err != nil {
		err := fmt.Errorf(`Invalid request body. err = %s`, err)
		code := 401
		gctx.AbortWithStatusJSON(code, err)
		return
	}
	agent.Coordinator = datatypes.AgentCoordinator{
		Name:        c.Name,
		LastCheckin: time.Now(),
		Status:      `healthy`,
	}

	c.Agents[agent.Name] = &agent

	gctx.JSON(200, gin.H{
		"status": "ok",
	})
}
