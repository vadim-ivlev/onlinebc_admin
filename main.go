package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"onlinebc_admin/controller"
	"onlinebc_admin/model/db"
	"onlinebc_admin/model/img"
	"onlinebc_admin/model/imgserver"
	"onlinebc_admin/model/redis"
	"onlinebc_admin/router"
	"os"
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
	db.CreateDatabaseIfNotExists()

	// печатаем конфиг базы данных для контроля ситуации
	if printParams {
		db.PrintConfig()
		return
	}

	// если servePort > 0, печатаем приветствие и запускаем сервер
	if servePort > 0 {
		greetings, _ := ioutil.ReadFile("./templates/greetings.txt")
		fmt.Printf(string(greetings), servePort)
		r := router.Setup(debug, true)
		r.Run(":" + strconv.Itoa(servePort))
	}
}

// Вспомогательные функции =========================================

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
		img.ReadConfig("./configs/img-docker.yaml")
	} else {
		db.ReadConfig("./configs/db-dev.yaml")
		redis.ReadConfig("./configs/redis-dev.yaml")
		imgserver.ReadConfig("./configs/imgserver-dev.yaml")
		img.ReadConfig("./configs/img-dev.yaml")
	}

	db.ReadConfig("./configs/db.yaml")
	redis.ReadConfig("./configs/redis.yaml")
	imgserver.ReadConfig("./configs/imgserver.yaml")
	img.ReadConfig("./configs/img.yaml")

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
