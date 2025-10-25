package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura para o corpo da requisição POST /api/find-service
type FindServiceRequest struct {
	Intent string `json:"Intent"`
}

// Estrutura de resposta padrão
type FindServiceResponse struct {
	Success bool        `json:"sucess"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// GET /api/healthz
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// POST /api/find-service
func FindService(c *gin.Context) {
	var req FindServiceRequest

	// Tenta decodificar o JSON recebido
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, FindServiceResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	// Por enquanto, uma resposta simulada (placeholder)
	c.JSON(http.StatusOK, FindServiceResponse{
		Success: true,
		Data: gin.H{
			"service_id":   1,
			"service_name": "Consulta Limite",
		},
		Error: "",
	})

}
