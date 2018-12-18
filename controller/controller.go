package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"onlinebc_admin/model/db"
	"strings"

	"github.com/gorilla/mux"
)

// Param - параметр запроса ?name=value&...
type Param struct {
	Comment string
	Name    string
	Value   string
}

// Route - маршрут.
type Route struct {
	Methods []string
	Comment string
	Path    string
	Example string
	Func    func(w http.ResponseWriter, r *http.Request) `json:"-" yaml:"-"`
	Params  []Param                                      `json:",omitempty" yaml:",omitempty"`
}

// ******************************************************************************************************

// Routes содержит инфориацию о маршрутах.  Документация API.
var Routes []Route

// GetMedium возвращает медиа по id
func GetMedium(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_medium($1);")
}

// GetPost возвращает пост по id
func GetPost(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_post($1);")
}

// GetBroadcast возвращает трансляцию по id
func GetBroadcast(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_broadcast($1);")
}

// UpdateMedium обновляет медиа по id
func UpdateMedium(w http.ResponseWriter, r *http.Request) {
	sqlText, vals := getUpdateSQL(mux.Vars(r))
	db.ExequteSQL(sqlText, vals...)
	// sqlText2 := getUpdateSQL2(mux.Vars(r))
	// db.ExequteSQL(sqlText2)
}

func getUpdateSQL(vars map[string]string) (string, []interface{}) {
	keys, values, qs := getKeysAndValues(vars)
	sqlText := fmt.Sprintf("UPDATE media SET (%s) = (%s) WHERE id = %s",
		strings.Join(keys, ", "),
		strings.Join(qs, ", "),
		vars["id"])

	v := make([]interface{}, 0)
	for _, s := range values {
		v = append(v, s)
	}

	return sqlText, v
}

func getUpdateSQL2(vars map[string]string) string {
	keys, values, _ := getKeysAndValues(vars)
	sqlText := fmt.Sprintf("UPDATE media SET (%s) = (%s) WHERE id = %s",
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
		vars["id"])
	return sqlText
}

// UpdatePost обновляет пост по id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_post($1);")
}

// UpdateBroadcast обновляет трансляцию по id
func UpdateBroadcast(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_broadcast($1);")
}

// ************************************************************************

// LandingPage : To test API in browser.
func LandingPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("templates/landing-page.html")
	if err == nil {
		tmpl.Execute(w, Routes)
	} else {
		fmt.Fprintf(w, "ERR=%v", err)
	}
}

// GetRoutes : Перечисляет доступные маршруты.  Документация API.
func GetRoutes(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.Marshal(Routes)
	fmt.Fprint(w, string(bytes))
}

// GetMedia возвращает все медиа поста по его id
func GetMedia(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_media($1);")
}

// GetAnswers возвращает ответы к посту по его id
func GetAnswers(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_answers($1);")
}

// GetPosts возвращает посты трансляции по её id
func GetPosts(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_posts($1);")
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func GetFullBroadcast(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_full_broadcast($1);")
}

// GetBroadcasts Получить список трансляций
func GetBroadcasts(w http.ResponseWriter, r *http.Request) {
	json := db.GetJSON("SELECT get_broadcasts();")
	fmt.Fprint(w, json)
}

// GetBroadcastList Получить список трансляций. Legacy.
func GetBroadcastList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	main := vars["main"]
	active := vars["active"]
	num := vars["num"]
	fmt.Printf("main=%v active=%v num=%v", main, active, num)
	json := db.GetJSON("SELECT get_broadcasts();")
	fmt.Fprint(w, json)
}

// FUNCTIONS *******************************************************

// getKeysAndValues возвращает срезы ключей и значений
func getKeysAndValues(m map[string]string) ([]string, []string, []string) {
	// l := len(m)
	keys := []string{}
	values := make([]string, 0)
	qustionMarks := []string{}

	for key, val := range m {
		keys = append(keys, key)
		values = append(values, "'"+val+"'")
		qustionMarks = append(qustionMarks, "?")
	}
	return keys, values, qustionMarks
}

// getByID возвращает что-то по его id в HTTP запросе запросе вида ?id=354
func getByID(w http.ResponseWriter, r *http.Request, sqlText string) {
	json := db.GetJSON(sqlText, mux.Vars(r)["id"])
	fmt.Fprint(w, json)
}
