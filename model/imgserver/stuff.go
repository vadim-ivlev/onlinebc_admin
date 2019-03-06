package imgserver

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Вспомогательные функции /////////////////////////////////////////////////////

type connectionParams struct {
	Host      string `yaml:"host"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Port      string `yaml:"port"`
	Uploaddir string `yaml:"uploaddir"`
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
