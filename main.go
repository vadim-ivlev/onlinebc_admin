package main

import (
	"onlinebc_admin/model/db"
	"onlinebc_admin/model/redis"
	"onlinebc_admin/router"
	"strconv"
)

func main() {
	// считать параметры командной строки
	servePort, front, debug, printParams := readCommandLineParams()
	// читаем конфиги Postgres, Redis и роутера.
	readConfigs(front)

	// Инициализируем Redis
	redis.Init()

	// Ждем готовности базы данных
	db.WaitForDbOrExit(10)

	// порождаем базу данных если ее нет
	createDatabaseIfNotExists()

	if printParams {
		db.PrintConfig()
		return
	}

	// если servePort > 0, печатаем приветствие и зпускаем сервер
	if servePort > 0 {
		printGreetings(servePort)
		r := router.Setup(debug, true)
		r.Run(":" + strconv.Itoa(servePort))
	}
}
