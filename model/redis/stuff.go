package redis

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Вспомогательные функции /////////////////////////////////////////////////////

type connectionParams struct {
	ConnectStr string `yaml:"connection string"`
	TTL        int    `yaml:"time to live"`
}

// Conf общие конфигурационные параметры
var params connectionParams

// ReadConfig читает файл YAML
func ReadConfig(fileName string, env string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err.Error())
		return
	}

	envParams := make(map[string]connectionParams)
	err = yaml.Unmarshal(yamlFile, &envParams)
	if err != nil {
		log.Println(err.Error())
	}
	params = envParams[env]

}
