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
	"strconv"
)

func main() {
	// считать параметры командной строки
	servePort, front, debug, printParams, env := readCommandLineParams()

	// читаем конфиги Postgres, Redis и роутера.
	readConfigs(front, env)

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

// readConfigs считывает конфиги Postgres, Redis, imgserver и img.
// frontEndMode - Пути и соответственно доступный API зависят от моды,
// в котрой запущено приложение: Фронтэнд или Бэкэнд.
// env - Окружение. Возможные значения: dev - разработка, docker - в докере для фронтэнд разработчиков. prod - по умолчанию для продакшн.
func readConfigs(frontEndMode bool, env string) {

	db.ReadConfig("./configs/db.yaml", env)
	redis.ReadConfig("./configs/redis.yaml", env)
	imgserver.ReadConfig("./configs/imgserver.yaml", env)
	img.ReadConfig("./configs/img.yaml", env)

	// Если запущен во Фронтэнд моде загрузить соответствующие руты.
	if frontEndMode {
		controller.ReadConfig("./configs/routes-front.yaml")
	} else {
		controller.ReadConfig("./configs/routes.yaml")
	}
}

// readCommandLineParams читает параметры командной строки
func readCommandLineParams() (serverPort int, front bool, debug bool, printParams bool, env string) {
	flag.IntVar(&serverPort, "serve", 0, "Запустить приложение на порту с номером > 0 ")
	flag.BoolVar(&front, "front", false, "Запустить приложение в режиме Фронтэнд. Рауты только на чтение.")
	flag.BoolVar(&debug, "debug", false, "Режим Debug. С отображением запросов в консоль.")
	flag.BoolVar(&printParams, "showparams", false, "Показать параметры соединения с БД.")
	flag.StringVar(&env, "env", "prod", "Окружение. Возможные значения: dev - разработка, docker - в докере для фронтэнд разработчиков. prod - по умолчанию для продакшн.")
	flag.Parse()
	fmt.Println("\nПример запуска: go build && ./onlinebc_admin -serve 7777")
	flag.Usage()
	return
}
