package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"onlinebc/model/db"
	"onlinebc/model/redis"
	"onlinebc/router"
	"os"
	"strconv"
)

func main() {
	// считать конфиги
	db.ReadConfig("./configs/db.yaml")
	redis.ReadConfig("./configs/redis.yaml")

	// считать параметры командной строки
	serve, port := readCommandLineParams()

	// если есть параметр -serve, напечатать приветствие и запустить сервер
	if serve {
		printGreetings(port)
		router.Serve(":" + strconv.Itoa(port))
	}
}

func readCommandLineParams() (bool, int) {
	port := 1234
	serve := true

	flag.IntVar(&port, "port", 7777, "Номер порта")
	flag.BoolVar(&serve, "serve", false, "Запустить приложение")

	initdb := flag.Bool("initdb", false, "Инициализировать БД c тестовыми данными.")
	createDbFunctions := flag.Bool("create-db-functions", false, "Породить функции БД из файла migrations/views-and-functions.sql")
	printParams := flag.Bool("print-params", false, "Показать параметры приложения.")

	flag.Parse()

	if *initdb {
		fmt.Println("инициализация БД...")
		db.ExequteSQL(readTextFile("./migrations/onlinebc-dump.sql"))
		os.Exit(0)
	}
	if *createDbFunctions {
		fmt.Println("Порождение функций БД...")
		db.ExequteSQL(readTextFile("./migrations/views-and-functions.sql"))
		os.Exit(0)
	}
	if *printParams {
		db.PrintConfig()
		redis.PrintConfig()
		os.Exit(0)
	}

	fmt.Println("\nПараметры запуска приложения **************************\n")
	flag.Usage()

	return serve, port
}

func printGreetings(port int) {
	fmt.Printf(readTextFile("./docs/greetings.txt"), port, port)
}

func readTextFile(fileName string) string {
	txt, _ := ioutil.ReadFile(fileName)
	return string(txt)
}
