package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GinHeadersMiddleware добавляет HTTP заголовки к ответу сервера
func GinHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Content-Type", "application/json; charset=utf-8")
		fmt.Println("header")
		c.Next()
	}
}
