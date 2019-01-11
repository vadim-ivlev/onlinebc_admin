package middleware

import (
	"bufio"
	"bytes"
	"onlinebc_admin/model/redis"
	"strings"

	"github.com/gin-gonic/gin"
)

// HeadersMiddleware добавляет HTTP заголовки к ответу сервера
func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println("**************HEADER****************")
		// c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}

// RedisMiddleware Проверяет нет ли закэшированного значения в Redis.
// Если есть посылает его клиенту.
func RedisMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.RequestURI
		if strings.HasPrefix(key, "/api/") {
			value, err := redis.Get(key)
			if err == nil {
				c.Header("Redis", "********** FROM REDIS CACHE ************")
				c.String(200, value)
				c.Abort()
				return
			}

			w := bufio.NewWriter(c.Writer)
			buff := bytes.Buffer{}
			newWriter := &bufferedWriter{c.Writer, w, buff}
			c.Writer = newWriter

			c.Next()

			if c.Writer.Status() == 200 {
				s := newWriter.Buffer.String()
				redis.Set(key, s)
			}
			// You have to manually flush the buffer at the end
			w.Flush()

		}

	}
}

type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}
