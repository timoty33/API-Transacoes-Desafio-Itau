package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"transactionsApi/handler"
	"transactionsApi/middleware"
	sot "transactionsApi/suportObsTest"
)

func Routes() *gin.Engine {
	fmt.Println("[server] Configurando as rotas...")

	server := gin.Default()
	server.Use(middleware.ErrorHandler()) // Use o middleware de erro

	server.POST("/transacoes", handler.SalvarTransacao)
	server.DELETE("/transacao/:senha", handler.DeletarTransacoes)
	server.GET("/estatistica/:duracao", handler.Estatisticas)
	server.GET("/health", sot.Health(server))
	fmt.Println("[server] Rotas criadas com sucesso!")

	return server
}
