package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"onlinebc/model/db"
	"onlinebc/model/redis"

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
	Comment string
	Path    string
	Example string
	Func    func(w http.ResponseWriter, r *http.Request) `json:"-" yaml:"-"`
	Params  []Param                                      `json:",omitempty" yaml:",omitempty"`
}

// Query строит строку HTTP запроса по параметрам. Применяется в шаблонах документации API.
func (r Route) Query() string {
	s := r.Path + "?"
	for _, p := range r.Params {
		s += p.Name + "=" + p.Value + "&"
	}
	return s[0 : len(s)-1]
}

// ******************************************************************************************************

// Routes содержит инфориацию о маршрутах.  Документация API.
var Routes []Route

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

// getByID возвращает что-то по его id в HTTP запросе запросе вида ?id=354
func getByID(w http.ResponseWriter, r *http.Request, sqlText string) {
	id := mux.Vars(r)["id"]
	json := db.GetJSON(sqlText, id)
	redis.Set(r.RequestURI, json)
	fmt.Fprint(w, json)
}

// GetMedia возвращает медиа поста по его id
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

// GetBroadcast возвращает трасляцию с постами по её id
func GetBroadcast(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_broadcast($1);")
}

// GetBroadcasts Получить список трансляций
func GetBroadcasts(w http.ResponseWriter, r *http.Request) {
	json := db.GetJSON("SELECT get_broadcasts();")
	redis.Set(r.RequestURI, json)
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
	redis.Set(r.RequestURI, json)
	fmt.Fprint(w, json)
}
