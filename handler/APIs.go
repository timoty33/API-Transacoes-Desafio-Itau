package handler

import (
	"fmt"
	"net/http"

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

func Estaticas(c *gin.Context) {
	fmt.Println("[handler] Coletando estatísticas...")

	ultimasTransacoes := utils.UltimasTransacoes(data.Transacoes)

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

	fmt.Println("[handler] Estatísticas enviadas!")
}

func Health(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[heath] Verificando rotas existentes...")

		rotas := engine.Routes()
		rotasMap := make(map[string]bool)

		for _, rotaAtual := range rotas {
			rotasMap[rotaAtual.Path] = true
		}

		rotasEsperadas := []string{"/transacoes", "/transacao/:senha", "/estatistica", "/health"}
		rotasFaltando := []string{}

		for _, rotaEsperada := range rotasEsperadas {
			if !rotasMap[rotaEsperada] {
				rotasFaltando = append(rotasFaltando, rotaEsperada)
			}
		}

		if len(rotasFaltando) != 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":        "Not OK",
				"message":       "Erro ao obter todas as rotas!",
				"rotasFaltando": rotasFaltando,
			})
			fmt.Println("[health] Erro nas Rotas!")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Nenhum problema encontrado, todas as rotas estão certas",
		})
		fmt.Println("[health] Rotas estão ajustadas")
	}
}
