package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"transactionsApi/handler"
	sot "transactionsApi/suportObsTest"
)

func Routes() *gin.Engine {
	fmt.Println("[server] Configurando as rotas...")

	server := gin.Default()

	server.POST("/transacoes", handler.SalvarTransacao)
	server.DELETE("/transacao/:senha", handler.DeletarTransacoes)
	server.GET("/estatistica/:duracao", handler.Estatisticas)
	server.GET("/health", sot.Health(server))
	fmt.Println("[server] Rotas criadas com sucesso!")

	return server
}
