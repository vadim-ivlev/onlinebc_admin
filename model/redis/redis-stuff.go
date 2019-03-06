package redis

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Вспомогательные функции /////////////////////////////////////////////////////

type connectionParams struct {
	ConnectStr string `yaml:"connection string"`
	TTL        int    `yaml:"time to live"`
}

// Conf Common config params
var params connectionParams

// ReadConfig reads YAML file
func ReadConfig(fileName string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = yaml.Unmarshal(yamlFile, &params)
	if err != nil {
		fmt.Println(err.Error())
	}
}
