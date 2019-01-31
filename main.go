package main

import (
	"onlinebc_admin/model/db"
	"onlinebc_admin/router"
	"strconv"
)

func main() {
	// считать параметры командной строки
	servePort, front, debug, printParams := readCommandLineParams()
	// читаем конфиги Postgres, Redis и роутера.
	readConfigs(front)

	db.ExitIfNoDB()
	// порождаем базу данных если ее нет
	createDatabaseIfNotExists()

	if printParams {
		db.PrintConfig()
		return
	}

	// если servePort > 0, печатаем приветствие и зпускаем сервер
	if servePort > 0 {
		printGreetings(servePort)
		// router.Serve(":"+strconv.Itoa(servePort), debug)
		r := router.SetupRouter(debug)
		r.Run(":" + strconv.Itoa(servePort))

	}
}
