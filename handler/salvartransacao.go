package handler

import (
	"fmt"
	"net/http"
	"transactionsApi/model"
	"transactionsApi/utils"

	"github.com/gin-gonic/gin"
)

func SalvarTransacao(c *gin.Context) {
	fmt.Println("[handler] Salvando transação...")

	var novaTransacao model.Transacao

	// Esse condicional além de captar se tem o JSON,
	// Também vê se todos os campos foram povoados
	if c.ShouldBindJSON(&novaTransacao) != nil {
		// BadRequest caso dê algum erro com o JSON recebido
		c.JSON(http.StatusBadRequest, gin.H{})
		fmt.Println("Erro de BadRequest")
		return
	}

	if !utils.ValidarTransacao(novaTransacao.DataHora, novaTransacao.Valor) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		fmt.Println("Erro de UnprocessableEntity")
		return
	}

	model.Transacoes = append(model.Transacoes, novaTransacao)

	// Retorna um status 201 (Created) e a transação que foi salva.
	c.JSON(http.StatusCreated, gin.H{})

	fmt.Println("[handler] Transação salva com sucesso!")

	fmt.Println(model.Transacoes)
}
