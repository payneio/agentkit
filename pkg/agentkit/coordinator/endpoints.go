package coordinator

import "github.com/gin-gonic/gin"

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
	gctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func (c *Coordinator) UpsertAgent(gctx *gin.Context) {
	gctx.JSON(200, gin.H{
		"status": "ok",
	})
}
