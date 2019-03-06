package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"onlinebc_admin/controller"
	"onlinebc_admin/model/db"
	"onlinebc_admin/model/imgserver"
	"onlinebc_admin/model/redis"

	"os"
)

// readConfigs считывает конфиги Postgres и Redis.
// Следующий конфиг перегружает предыдущий,
// так что может присутствовать несколько конфигов для db и Redis.
// Пути и соответственно доступный API зависят от моды,
// в котрой запущено приложение: Фронтэнд или Бэкэнд.
func readConfigs(frontEndMode bool) {
	if os.Getenv("RUNNING_IN_DOCKER") == "Y" {
		db.ReadConfig("./configs/db-docker.yaml")
		redis.ReadConfig("./configs/redis-docker.yaml")
		imgserver.ReadConfig("./configs/imgserver-docker.yaml")
	} else {
		db.ReadConfig("./configs/db-dev.yaml")
		redis.ReadConfig("./configs/redis-dev.yaml")
		imgserver.ReadConfig("./configs/imgserver-dev.yaml")
	}

	db.ReadConfig("./configs/db.yaml")
	redis.ReadConfig("./configs/redis.yaml")
	imgserver.ReadConfig("./configs/imgserver.yaml")

	// Если запущен во Фронтэнд моде загрузить соответствующие руты.
	if frontEndMode {
		controller.ReadConfig("./configs/routes-front.yaml")
	} else {
		controller.ReadConfig("./configs/routes.yaml")
	}
}

// readCommandLineParams читает параметры командной строки
func readCommandLineParams() (serverPort int, front bool, debug bool, printParams bool) {
	flag.IntVar(&serverPort, "serve", 0, "Запустить приложение на порту с номером > 0 ")
	flag.BoolVar(&front, "front", false, "Запустить приложение в режиме Фронтэнд. Рауты только на чтение.")
	flag.BoolVar(&debug, "debug", false, "Режим Debug. С отображением запросов в консоль.")
	flag.BoolVar(&printParams, "showparams", false, "Показать параметры соединения с БД.")
	flag.Parse()
	fmt.Println("\nПример запуска: go build && ./onlinebc_admin -serve 7777")
	flag.Usage()
	return
}

// printGreetings печатает красивое сообщение после запуска программы
func printGreetings(port int) {
	fmt.Printf(getTextFromFile("./templates/greetings.txt"), port)
}

// getTextFromFile возвращает текст файла
func getTextFromFile(fileName string) string {
	txt, _ := ioutil.ReadFile(fileName)
	return string(txt)
}

func createDatabaseIfNotExists() {
	fmt.Println("Порождение таблиц ...")
	db.GetExecResult(getTextFromFile("./migrations/create-tables.sql"))
	fmt.Println("Порождение функций ...")
	db.GetExecResult(getTextFromFile("./migrations/create-views-and-functions.sql"))
	fmt.Println("Наполнение тестовыми данными...")
	db.GetExecResult(getTextFromFile("./migrations/add-data.sql"))
}
