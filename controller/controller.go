package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"onlinebc_admin/model/db"

	"github.com/gorilla/mux"
)

// ************************************************************************

// GetMedium возвращает медиа по id
func (dummy) GetMedium(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_medium($1);")
}

// GetPost возвращает пост по id
func (dummy) GetPost(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_post($1);")
}

// GetBroadcast возвращает трансляцию по id
func (dummy) GetBroadcast(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_broadcast($1);")
}

// ************************************************************************

// UpdateMedium обновляет медиа по id
func (dummy) UpdateMedium(w http.ResponseWriter, r *http.Request) {
	db.UpdateRowByID("media", mux.Vars(r))
}

// UpdatePost обновляет пост по id
func (dummy) UpdatePost(w http.ResponseWriter, r *http.Request) {
	db.UpdateRowByID("post", mux.Vars(r))
}

// UpdateBroadcast обновляет трансляцию по id
func (dummy) UpdateBroadcast(w http.ResponseWriter, r *http.Request) {
	db.UpdateRowByID("broadcast", mux.Vars(r))
}

// ************************************************************************

// CreateMedium обновляет медиа по id
func (dummy) CreateMedium(w http.ResponseWriter, r *http.Request) {
	db.CreateRow("media", mux.Vars(r))
}

// CreatePost обновляет пост по id
func (dummy) CreatePost(w http.ResponseWriter, r *http.Request) {
	db.CreateRow("post", mux.Vars(r))
}

// CreateBroadcast обновляет трансляцию по id
func (dummy) CreateBroadcast(w http.ResponseWriter, r *http.Request) {
	db.CreateRow("broadcast", mux.Vars(r))
}

// ************************************************************************

// DeleteMedium обновляет медиа по id
func (dummy) DeleteMedium(w http.ResponseWriter, r *http.Request) {
	db.DeleteRowByID("media", mux.Vars(r))
}

// DeletePost обновляет пост по id
func (dummy) DeletePost(w http.ResponseWriter, r *http.Request) {
	db.DeleteRowByID("post", mux.Vars(r))
}

// DeleteBroadcast обновляет трансляцию по id
func (dummy) DeleteBroadcast(w http.ResponseWriter, r *http.Request) {
	db.DeleteRowByID("broadcast", mux.Vars(r))
}

// ************************************************************************

// LandingPage : To test API in browser.
func (dummy) LandingPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("templates/landing-page.html")
	if err == nil {
		tmpl.Execute(w, Routes)
	} else {
		fmt.Fprintf(w, "ERR=%v", err)
	}
}

// GetRoutes : Перечисляет доступные маршруты.  Документация API.
func (dummy) GetRoutes(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.Marshal(Routes)
	// bytes, _ := yaml.Marshal(Routes)
	fmt.Fprint(w, string(bytes))
}

// ************************************************************************

// GetMedia возвращает все медиа поста по его id
func (dummy) GetMedia(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_media($1);")
}

// GetAnswers возвращает ответы к посту по его id
func (dummy) GetAnswers(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_answers($1);")
}

// GetPosts возвращает посты трансляции по её id
func (dummy) GetPosts(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_posts($1);")
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func (dummy) GetFullBroadcast(w http.ResponseWriter, r *http.Request) {
	getByID(w, r, "SELECT get_full_broadcast($1);")
}

// GetBroadcasts Получить список трансляций
func (dummy) GetBroadcasts(w http.ResponseWriter, r *http.Request) {
	json := db.GetJSON("SELECT get_broadcasts();")
	fmt.Fprint(w, json)
}

// GetBroadcastList Получить список трансляций. Legacy.
func (dummy) GetBroadcastList(w http.ResponseWriter, r *http.Request) {
	json := db.GetJSON("SELECT get_broadcasts();")
	fmt.Fprint(w, json)
}
