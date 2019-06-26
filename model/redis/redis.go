package redis

import (
	"fmt"
	"log"
	"onlinebc_admin/model/db"
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
	fmt.Println("Redis Get:", key)
	return redisdb.Get("onlinebc:" + key).Result()
}

// Set сохраняет значение ключа в Redis на установленное время.
func Set(key string, value string) error {
	fmt.Println("Redis Set:", key)
	return redisdb.Set("onlinebc:"+key, value, time.Second*time.Duration(params.TTL)).Err()
}

// Del удаляет запись по ключу
func Del(key string) error {
	fmt.Println("Redis Del:", key)
	return redisdb.Del("onlinebc:" + key).Err()
}

// I N V A L I D A T I O N  ======================================================================

// ClearByBroadcastID чистим redis по id трансляции
func ClearByBroadcastID(id interface{}) {
	if id == nil {
		log.Println("ClearByBroadcastID: no id")
		return
	}
	key := fmt.Sprintf("/api/full-broadcast/%v?", id)
	err := Del(key)
	if err != nil {
		log.Println(err)
	}
}

// ClearByPostID чистим redis по id поста.
func ClearByPostID(id interface{}) {
	if id == nil {
		log.Println("ClearByPostID: no id")
		return
	}

	// получаем идентификатор трансляции и идентификатор родителя на случай если это ответ к посту
	id_broadcast, id_parent, err := GetPostParentIDs(id)
	if err != nil {
		log.Println("ClearByPostID: SELECT: ", err)
		return
	}

	// если это пост чистим трансляцию
	if id_broadcast.(int64) > 0 {
		ClearByBroadcastID(id_broadcast)
		return
	}

	// если это ответ к посту рекурсивно вызываем себя с параметром id_parent
	if id_parent.(int64) > 0 {
		ClearByPostID(id_parent)
		return
	}
}

// ClearByImageID чистим redis по id изображения.
func ClearByImageID(id interface{}) {
	if id == nil {
		log.Println("ClearByImageID: no id")
		return
	}

	// получаем идентификатор поста
	post_id := GetImagePostID(id)
	// чистим трансляцию по идентификатору поста
	ClearByPostID(post_id)
}

// GetPostParentIDs возвращает идентификатор трансляции и идентификатор родителя поста
func GetPostParentIDs(id interface{}) (interface{}, interface{}, error) {
	row, err := db.QueryRowMap("SELECT id_broadcast, id_parent FROM post WHERE id = $1 ;", id)
	if err != nil {
		log.Println("ClearByPostID: SELECT: ", err)
		return nil, nil, err
	}
	return row["id_broadcast"], row["id_parent"], nil
}

// GetImagePostID возвращает идентификатор поста изображения
func GetImagePostID(id interface{}) interface{} {
	row, err := db.QueryRowMap("SELECT post_id FROM image WHERE id = $1 ;", id)
	if err != nil {
		log.Println("GetImagePostID:", err)
		return nil
	}
	return row["post_id"]
}
