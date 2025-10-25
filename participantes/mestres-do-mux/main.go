package mestresdomux

import (
	"github.com/gin-gonic/gin"
	"routes"
	"os"
)

func main (){
	//lê variavel de ambiente PORT, com fallback para 18020
	port := os.Getenv("PORT")
	if port == ""{
		port = "18020"
	}

	//cria o servidor GIng com middleware padrão (logger e recovery)
	r := gin.Default()

	//Inicializa as rotas
	routes.RegisterRoutes(r)

	//roda o servidor
	r.Run(":" + port)
}

/*
	🧠 Explicando:

	gin.Default() → cria uma instância do servidor com logs e tratamento de erros automático

	routes.RegisterRoutes(r) → importa as rotas de outro arquivo (mantém o código organizado)

	r.Run(":18020") → inicia o servidor HTTP na porta definida

*/