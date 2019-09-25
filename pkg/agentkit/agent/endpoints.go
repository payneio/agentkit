package agent

import (
	"agentkit/pkg/agentkit/datatypes"

	"github.com/gin-gonic/gin"
)

func (a *Agent) WebRoot(gctx *gin.Context) {
	response := struct {
		Name          string            `json:"name"`
		Port          int               `json:"port"`
		PublicAddress string            `json:"publicAddress"`
		Central       datatypes.Central `json:"central"`
	}{
		Name:          a.Name,
		Port:          a.Port,
		PublicAddress: a.PublicAddress,
		Central:       a.Central,
	}
	gctx.JSON(200, response)
}

func (a *Agent) WebHealth(gctx *gin.Context) {
	gctx.JSON(200, gin.H{"status": "ok"})
}

func (a *Agent) WebReadMind(gctx *gin.Context) {
	gctx.JSON(200, a.Mind.GetBeliefs().MSI())
}
