package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"transactionsApi/data"
	"transactionsApi/utils"
)

func SalvarTransacao(c *gin.Context) {
	fmt.Println("[handler] Salvando transação...")

	var novaTransacao data.Transacao

	// Esse condicional além de captar se tem o JSON,
	// Também vê se todos os campos foram povoados
	if err := c.ShouldBindJSON(&novaTransacao); err != nil {
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

	data.Transacoes = append(data.Transacoes, novaTransacao)

	// Retorna um status 201 (Created) e a transação que foi salva.
	c.JSON(http.StatusCreated, gin.H{})

	fmt.Println("[handler] Transação salva com sucesso!")

	fmt.Println(data.Transacoes)
}

func DeletarTransacoes(c *gin.Context) {
	fmt.Println("[handler] Deletando transações...")

	senhaCerta := "Timoteo2011@"
	senhaObtida := c.Param("senha")

	if senhaObtida == senhaCerta {
		data.Transacoes = make([]data.Transacao, 0)

		c.JSON(http.StatusOK, gin.H{})

		fmt.Println("[handler] Transações deletadas com sucesso!")
		return
	}

	fmt.Println("[handler] Senha recebida está incorreta!")
	c.JSON(http.StatusBadRequest, gin.H{})
}

func Estatisticas(c *gin.Context) {
	fmt.Println("[handler] Coletando estatísticas...")

	duracaoStr := c.Param("duracao")
	duracao, err := time.ParseDuration(duracaoStr + "s")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Duração em um formato invalido (exemplo de como precisa ser: 120, para 120 segundos"})
		return
	}

	ultimasTransacoes := utils.UltimasTransacoes(data.Transacoes, duracao) // transações que aconteceram nos últimos X segundos

	count := len(ultimasTransacoes)
	sum := utils.Soma(ultimasTransacoes)
	avg := utils.Media(ultimasTransacoes)
	min, max := utils.MinMax(ultimasTransacoes)

	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"sum":   sum,
		"avg":   avg,
		"min":   min,
		"max":   max,
	})

	fmt.Println("[handler] Estatísticas enviadas com uma janela de duração de:", duracaoStr)
}
