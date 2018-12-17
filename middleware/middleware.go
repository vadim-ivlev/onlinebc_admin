package middleware

import (
	"net/http"
)

// HeadersMiddleware добавляет CORS заголовки к ответу сервера
// для кроссдоменных запросов.
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
