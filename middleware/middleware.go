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

		// url откуда пришел запрос
		// protocol := "https://"
		// if strings.HasPrefix(c.Request.Proto, "HTTP/") {
		// 	protocol = "http://"
		// }
		// host := protocol + c.Request.Host

		c.Header("Access-Control-Allow-Origin", "https://editor.rg.ru")
		if hostIsAllowed(c.Request.Host) {
			c.Header("Access-Control-Allow-Origin", "*")
			// c.Header("Access-Control-Allow-Origin", host)
		}
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Max-Age", "600")
		c.Next()
	}
}

// hostIsAllowed - проверяем пришел ли запрос с разрешенного хоста
func hostIsAllowed(host string) bool {
	if strings.HasPrefix(host, "localhost") ||
		strings.HasPrefix(host, "127.0.0.1") ||
		strings.Contains(host, ".rg.ru:") ||
		strings.HasSuffix(host, ".rg.ru") {
		return true
	}
	return false
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
				// break the chain
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
			// manually flush the buffer at the end
			w.Flush()

		}

	}
}

// bufferedWriter используется вместо ResposeWriter чтобы
// постфактум читать что было записано в поток. Переменная Buffer служит для этой цели.
type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}
