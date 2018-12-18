package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"onlinebc_admin/model/db"

	"github.com/gorilla/mux"
	yaml "gopkg.in/yaml.v2"
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
	db.UpdateRowByID("media", mux.Vars(r))
}

// UpdatePost обновляет пост по id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db.UpdateRowByID("post", mux.Vars(r))
}

// UpdateBroadcast обновляет трансляцию по id
func UpdateBroadcast(w http.ResponseWriter, r *http.Request) {
	db.UpdateRowByID("broadcast", mux.Vars(r))
}

// CreateMedium обновляет медиа по id
func CreateMedium(w http.ResponseWriter, r *http.Request) {
	db.CreateRow("media", mux.Vars(r))
}

// CreatePost обновляет пост по id
func CreatePost(w http.ResponseWriter, r *http.Request) {
	db.CreateRow("post", mux.Vars(r))
}

// CreateBroadcast обновляет трансляцию по id
func CreateBroadcast(w http.ResponseWriter, r *http.Request) {
	db.CreateRow("broadcast", mux.Vars(r))
}

// DeleteMedium обновляет медиа по id
func DeleteMedium(w http.ResponseWriter, r *http.Request) {
	db.DeleteRowByID("media", mux.Vars(r))
}

// DeletePost обновляет пост по id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	db.DeleteRowByID("post", mux.Vars(r))
}

// DeleteBroadcast обновляет трансляцию по id
func DeleteBroadcast(w http.ResponseWriter, r *http.Request) {
	db.DeleteRowByID("broadcast", mux.Vars(r))
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
	// bytes, _ := json.Marshal(Routes)
	bytes, _ := yaml.Marshal(Routes)
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

// getByID возвращает что-то по его id в HTTP запросе запросе вида ?id=354
func getByID(w http.ResponseWriter, r *http.Request, sqlText string) {
	json := db.GetJSON(sqlText, mux.Vars(r)["id"])
	fmt.Fprint(w, json)
}
