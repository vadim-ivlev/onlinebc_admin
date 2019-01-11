package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"onlinebc_admin/controller"
	"onlinebc_admin/model/db"
	"onlinebc_admin/model/redis"
	"onlinebc_admin/router"
	"os"
	"strconv"
)

func main() {

	// считать конфиги Postgres и Redis. Следующий конфиг перегружает предыдущий

	if os.Getenv("RUNNING_IN_DOCKER") == "Y" {
		db.ReadConfig("./configs/db-docker.yaml")
		redis.ReadConfig("./configs/redis-docker.yaml")
	} else {
		db.ReadConfig("./configs/db-dev.yaml")
		redis.ReadConfig("./configs/redis-dev.yaml")
	}

	db.ReadConfig("./configs/db.yaml")
	redis.ReadConfig("./configs/redis.yaml")

	createDatabaseWithData()

	// считать параметры командной строки
	servePort, front, debug := readCommandLineParams()

	// если serve > 0, напечатать приветствие и запустить сервер
	if servePort > 0 {
		printGreetings(servePort)

		// Если запущен во Фронтэнд моде загрузить соответствующие руты.
		if front {
			controller.ReadConfig("./configs/routes-front.yaml")
		} else {
			controller.ReadConfig("./configs/routes.yaml")
		}

		router.Serve(":"+strconv.Itoa(servePort), debug)
	}
}

func readCommandLineParams() (serverPort int, front bool, debug bool) {
	// serverPort = 0
	flag.IntVar(&serverPort, "serve", 0, "Запустить приложение на порту с номером > 0 ")

	// front = false
	flag.BoolVar(&front, "front", false, "Запустить приложение в режиме Фронтэнд. Рауты только на чтение.")

	// debug = false
	flag.BoolVar(&debug, "debug", false, "Режим Debug. С отображением запросов в консоль.")

	printParams := flag.Bool("showparams", false, "Показать параметры соединения с БД.")

	flag.Parse()

	if *printParams {
		db.PrintConfig()
		os.Exit(0)
	}

	fmt.Println("\nПример запуска: go run main.go -serve 7777 \n")
	flag.Usage()

	return
}

func printGreetings(port int) {
	fmt.Printf(readTextFile("./templates/greetings.txt"), port)
}

func readTextFile(fileName string) string {
	txt, _ := ioutil.ReadFile(fileName)
	return string(txt)
}

func createDatabaseWithData() {
	fmt.Println("Порождение таблиц ...")
	db.GetExecResult(readTextFile("./migrations/create-tables.sql"))
	fmt.Println("Порождение функций ...")
	db.GetExecResult(readTextFile("./migrations/create-views-and-functions.sql"))
	fmt.Println("Наполнение тестовыми данными...")
	db.GetExecResult(readTextFile("./migrations/add-data.sql"))
}
