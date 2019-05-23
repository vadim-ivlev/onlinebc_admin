package controller

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

// Param - параметр запроса ?name=value&...
type Param struct {
	Comment string
	Name    string
	Inptype string
	Value   string
	Test    string
}

// Route - маршрут.
type Route struct {
	Comment    string
	Methods    []string
	Path       string
	Example    string
	Controller string
	Params     []Param `json:",omitempty" yaml:",omitempty"`
}

type dummy struct {
}

// ************************************************************************

// Routes содержит информацию о маршрутах.  Документация API.
var Routes []Route

// GetFunctionByName возвращает функцию по имени
func GetFunctionByName(funcName string) func(*gin.Context) {
	m := reflect.ValueOf(&dummy{}).MethodByName(funcName)
	mCallable := m.Interface().(func(*gin.Context))
	return mCallable
}

// ReadConfig reads YAML file
func ReadConfig(fileName string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Routes)
	if err != nil {
		fmt.Println(err.Error())
	}
}
