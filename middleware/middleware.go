package middleware

import (
	"net/http"
)

// HeadersMiddleware добавляет CORS заголовки к ответу сервера
// для кроссдоменных запросов.
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// new
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		// w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		// w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		next.ServeHTTP(w, r)
	})
}
