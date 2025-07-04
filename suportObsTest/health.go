package suportObsTests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[health] Verificando rotas existentes...")

		rotas := engine.Routes()
		rotasMap := make(map[string]bool)

		for _, rotaAtual := range rotas {
			rotasMap[rotaAtual.Path] = true
		}

		rotasEsperadas := []string{"/transacoes", "/transacao/:senha", "/estatistica/:duracao", "/health"}
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
