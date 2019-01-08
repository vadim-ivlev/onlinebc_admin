package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
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

func getIntID(r *http.Request) int {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	return id
}

// GetFunctionByName возвращает функцию по имени
func GetFunctionByName(funcName string) func(http.ResponseWriter, *http.Request) {
	m := reflect.ValueOf(&dummy{}).MethodByName(funcName)
	mCallable := m.Interface().(func(http.ResponseWriter, *http.Request))
	return mCallable
}

// getFormFields извлекает хэш имен-значений полей формы из запроса
// or builds a map with keys "query", "variables", "operationName".
// Decoded body has precedence over POST over GET.
func getPayload(r *http.Request) map[string]interface{} {
	m := make(map[string]interface{})
	r.ParseForm()
	for k := range r.Form {
		m[k] = r.FormValue(k)
	}
	json.NewDecoder(r.Body).Decode(&m)
	return m
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
