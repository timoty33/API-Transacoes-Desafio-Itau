package handler

import (
	"fmt"
	"net/http"
	"time"
	"transactionsApi/model"
	"transactionsApi/utils"

	"github.com/gin-gonic/gin"
)

func Estatisticas(c *gin.Context) {
	fmt.Println("[handler] Coletando estatísticas...")

	duracaoStr := c.Param("duracao")
	duracao, err := time.ParseDuration(duracaoStr + "s")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Duração em um formato invalido (exemplo de como precisa ser: 120, para 120 segundos"})
		return
	}

	ultimasTransacoes := utils.UltimasTransacoes(model.Transacoes, duracao) // transações que aconteceram nos últimos X segundos

	count := len(ultimasTransacoes)
	sum := utils.Soma(ultimasTransacoes)
	avg, err := utils.Media(ultimasTransacoes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular a média: " + err.Error()})
		return
	}
	min, max, err := utils.MinMax(ultimasTransacoes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular min/max: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"sum":   sum,
		"avg":   avg,
		"min":   min,
		"max":   max,
	})

	fmt.Println("[handler] Estatísticas enviadas com uma janela de duração de:", duracaoStr)
}
