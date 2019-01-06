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
	json.NewEncoder(w).Encode(db.GetRowByID("medium", getIntID(r)))
}

// GetPost возвращает пост по id
func (dummy) GetPost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetRowByID("post", getIntID(r)))
}

// GetBroadcast возвращает трансляцию по id
func (dummy) GetBroadcast(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetRowByID("broadcast", getIntID(r)))
}

// ************************************************************************

// UpdateMedium обновляет медиа по id
func (dummy) UpdateMedium(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.UpdateRowByID("medium", getIntID(r), getFormFields(r)))
}

// UpdatePost обновляет пост по id
func (dummy) UpdatePost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.UpdateRowByID("post", getIntID(r), getFormFields(r)))
}

// UpdateBroadcast обновляет трансляцию по id
func (dummy) UpdateBroadcast(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.UpdateRowByID("broadcast", getIntID(r), getFormFields(r)))
}

// ************************************************************************

// CreateMedium Создать медиа поста с идентификатором id
func (dummy) CreateMedium(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.CreateRow("medium", getFormFields(r)))
}

// CreatePost Создать пост трансляции с идентификатором id
func (dummy) CreatePost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.CreateRow("post", getFormFields(r)))
}

// CreateBroadcast Создать  трансляцию
func (dummy) CreateBroadcast(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.CreateRow("broadcast", getFormFields(r)))
}

// ************************************************************************

// DeleteMedium обновляет медиа по id
func (dummy) DeleteMedium(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.DeleteRowByID("medium", getIntID(r)))
}

// DeletePost обновляет пост по id
func (dummy) DeletePost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.DeleteRowByID("post", getIntID(r)))
}

// DeleteBroadcast обновляет трансляцию по id
func (dummy) DeleteBroadcast(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.DeleteRowByID("broadcast", getIntID(r)))
}

// ************************************************************************

// LandingPage : To test API in browser.
func (dummy) LandingPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("templates/index.html", "templates/index.css")
	if err == nil {
		tmpl.Execute(w, Routes)
	} else {
		fmt.Fprintf(w, "ERR=%v", err)
	}
}

// GetRoutes : Перечисляет доступные маршруты.  Документация API.
func (dummy) GetRoutes(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.Marshal(Routes)
	fmt.Fprint(w, string(bytes))
}

// ************************************************************************

// GetMedia возвращает все медиа поста по его id
func (dummy) GetMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_media($1);", mux.Vars(r)["id"]))

}

// GetAnswers возвращает ответы к посту по его id
func (dummy) GetAnswers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_answers($1);", mux.Vars(r)["id"]))
}

// GetPosts возвращает посты трансляции по её id
func (dummy) GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_posts($1);", mux.Vars(r)["id"]))
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func (dummy) GetFullBroadcast(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_full_broadcast($1);", mux.Vars(r)["id"]))
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func (dummy) GetFullBroadcastLegacy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_full_broadcast($1);", id))
}

// GetBroadcasts Получить список трансляций
func (dummy) GetBroadcasts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_broadcasts();"))
}

// GetBroadcastList Получить список трансляций. Legacy.
func (dummy) GetBroadcastList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.QueryRowResult("SELECT get_broadcasts();"))
}
