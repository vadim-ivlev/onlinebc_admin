package redis

import (
	"log"

	// "github.com/go-redis/redis"
	redis0 "github.com/mediocregopher/radix.v2/redis"
)

// Get returns value by key.
// If key is absent returns an empty string and an error.
func Get(key string) (string, error) {
	conn, err := redis0.Dial("tcp", params.ConnectStr)
	if err != nil {
		log.Print("Get No connection", params.ConnectStr)
		return "", err
	}
	defer conn.Close()

	str, err := conn.Cmd("GET", "onlinebc:"+key).Str()
	if err != nil {
		return "", err
	}

	return str, nil
}

// Set sets value by key.
func Set(key string, value string) error {
	conn, err := redis0.Dial("tcp", params.ConnectStr)
	if err != nil {
		log.Print("Set No connection", params.ConnectStr)
		return err
	}
	defer conn.Close()

	resp := conn.Cmd("SETEX", "onlinebc:"+key, params.TTL, value)
	if resp.Err != nil {
		return resp.Err
	}

	return nil
}

// Del deletes the key from the redis0.
func Del(key string) error {
	conn, err := redis0.Dial("tcp", params.ConnectStr)
	if err != nil {
		log.Print("DEL No connection")
		return err
	}
	defer conn.Close()

	resp := conn.Cmd("DEL", "onlinebc:"+key)
	if resp.Err != nil {
		return resp.Err
	}

	return nil
}

// Get1 returns value by key.
// If key is absent returns an empty string and an error.
// func Get1(key string) (string, error) {

// 	client := redis.NewClient(&redis.Options{
// 		Addr: params.ConnectStr,
// 	})

// 	if err != nil {
// 		log.Print("Get No connection", params.ConnectStr)
// 		return "", err
// 	}
// 	defer conn.Close()

// 	str, err := conn.Cmd("GET", "onlinebc:"+key).Str()
// 	if err != nil {
// 		return "", err
// 	}

// 	return str, nil
// }
