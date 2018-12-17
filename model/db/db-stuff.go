package db

import (
	"fmt"
	"io/ioutil"

	//blank import
	_ "github.com/lib/pq"
	yaml "gopkg.in/yaml.v2"
)

// Вспомогательные функции /////////////////////////////////////////////////////

type connectionParams struct {
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

var params connectionParams
var connectStr string

// ReadConfig reads YAML file
func ReadConfig(fileName string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	panicIf(err)
	err = yaml.Unmarshal(yamlFile, &params)
	printIf(err)
	connectStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", params.User, params.Password, params.Dbname, params.Sslmode)
}

// PrintConfig prints DB connection parameters.
func PrintConfig() {
	s, _ := yaml.Marshal(params)
	fmt.Printf("\nDB connection parameters:\n%s\n", s)
	fmt.Printf("\nDB connection string: %s\n", connectStr)
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
