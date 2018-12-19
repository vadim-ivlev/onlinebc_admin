package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"onlinebc_admin/model/db"
	"reflect"

	"github.com/gorilla/mux"
	yaml "gopkg.in/yaml.v2"
)

// Param - параметр запроса ?name=value&...
type Param struct {
	Comment string
	Name    string
	Value   string
	Test    string
}

// Route - маршрут.
type Route struct {
	Comment string
	Methods []string
	Path    string
	Example string
	// Func       func(w http.ResponseWriter, r *http.Request) `json:"-" yaml:"-"`
	Controller string
	Params     []Param `json:",omitempty" yaml:",omitempty"`
}

type dummy struct {
}

// ************************************************************************

// Routes содержит информацию о маршрутах.  Документация API.
var Routes []Route

// FUNCTIONS *******************************************************

// GetFunctionByName возвращает функцию по имени
func GetFunctionByName(funcName string) func(http.ResponseWriter, *http.Request) {
	m := reflect.ValueOf(&dummy{}).MethodByName(funcName)
	mCallable := m.Interface().(func(http.ResponseWriter, *http.Request))
	return mCallable
}

// getByID возвращает что-то по его id в HTTP запросе запросе вида ?id=354
func getByID(w http.ResponseWriter, r *http.Request, sqlText string) {
	json := db.GetJSON(sqlText, mux.Vars(r)["id"])
	fmt.Fprint(w, json)
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
