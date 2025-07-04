package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Erro não tratado: %v\n", r) // Log do erro
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro interno."})
				c.Abort()
			}
		}()
		c.Next()
	}
}
