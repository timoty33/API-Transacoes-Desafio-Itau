package handler

import (
	"fmt"
	"net/http"
	"transactionsApi/model"

	"github.com/gin-gonic/gin"
)

func DeletarTransacoes(c *gin.Context) {
	fmt.Println("[handler] Deletando transações...")

	senhaCerta := "Timoteo2011@"
	senhaObtida := c.Param("senha")

	if senhaObtida == senhaCerta {
		model.Transacoes = make([]model.Transacao, 0)

		c.JSON(http.StatusOK, gin.H{})

		fmt.Println("[handler] Transações deletadas com sucesso!")
		return
	}

	fmt.Println("[handler] Senha recebida está incorreta!")
	c.JSON(http.StatusBadRequest, gin.H{})
}
