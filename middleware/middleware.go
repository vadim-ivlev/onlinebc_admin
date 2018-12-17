package middleware

import (
	"net/http"
	"onlinebc/model/redis"
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

// RedisMiddleware. Проверяет нет ли закэшированного значения в Redis. 
// И если есть посылает его клиенту, освобождая контроллер 
// от слишком частых обращений к базе данных.
func RedisMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.RequestURI
		value, err := redis.Get(key)
		if err == nil {
			w.Header().Set("Redis", "Data restored from redis")
			w.Write([]byte(value))
			return
		}

		next.ServeHTTP(w, r)
	})
}
