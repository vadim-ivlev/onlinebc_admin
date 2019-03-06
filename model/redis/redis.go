package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// пул соединений
var redisdb *redis.Client

// Init создает пул соединений и проверяет Redis
func Init() {
	redisdb = redis.NewClient(&redis.Options{Addr: params.ConnectStr})
	pong, err := redisdb.Ping().Result()
	fmt.Println("REDIS INIT:", pong, err)
}

// Get возвращает значение по ключу
func Get(key string) (string, error) {
	return redisdb.Get("onlinebc:" + key).Result()
}

// Set сохраняет значение ключа в Redis на установленное время.
func Set(key string, value string) error {
	return redisdb.Set("onlinebc:"+key, value, time.Second*time.Duration(params.TTL)).Err()
}

// Del удаляет запись по ключу
func Del(key string) error {
	return redisdb.Del("onlinebc:" + key).Err()
}
