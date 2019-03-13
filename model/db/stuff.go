package db

import (
	"fmt"
	"io/ioutil"

	// "onlinebc_admin/model/db"

	"github.com/golang-migrate/migrate"
	//blank import
	_ "github.com/golang-migrate/migrate/database/postgres"
	//blank import
	_ "github.com/golang-migrate/migrate/source/file"

	//blank import
	_ "github.com/lib/pq"
	yaml "gopkg.in/yaml.v2"
)

// Вспомогательные функции /////////////////////////////////////////////////////

type connectionParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

var params connectionParams
var connectStr string
var connectURL string

// ReadConfig reads YAML file
func ReadConfig(fileName string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = yaml.Unmarshal(yamlFile, &params)
	printIf(err)
	connectStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", params.Host, params.Port, params.User, params.Password, params.Dbname, params.Sslmode)
	connectURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", params.User, params.Password, params.Host, params.Port, params.Dbname, params.Sslmode)

}

// PrintConfig prints DB connection parameters.
func PrintConfig() {
	s, _ := yaml.Marshal(params)
	fmt.Printf("\nDB connection parameters:\n%s\n", s)
	fmt.Printf("DB connection string: %s\n", connectStr)
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func printIf(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// CreateDatabaseIfNotExists порождает объекты базы данных и наполняет базу тестовыми данными
func CreateDatabaseIfNotExists() {
	fmt.Println("Миграция ...")
	m, err := migrate.New("file://migrations/", connectURL)
	panicIf(err)
	printIf(m.Up())
}

// getTextFromFile возвращает текст файла
func getTextFromFile(fileName string) string {
	txt, _ := ioutil.ReadFile(fileName)
	return string(txt)
}
