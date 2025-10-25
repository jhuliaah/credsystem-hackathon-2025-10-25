package routes

import (
	"github.com/gin-gonic/gin"
	"hackathon-api/controllers"
)

func RegisterRoutes(r *gin.Engine)
{
	api := r.Group("/api")
	{
		api.GET("/healthz", controllers.HealthCheck)
		api.POST("/find-service", controllers.FindService)
	}
}

/*
	🧠 Explicando:
	r.Group("/api") → cria um grupo de rotas com o prefixo /api
	Cada rota chama uma função controladora (que criaremos a seguir)
*/

