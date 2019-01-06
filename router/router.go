package router

import (
	"log"
	"net/http"
	c "onlinebc_admin/controller"
	"onlinebc_admin/middleware"

	"github.com/gorilla/mux"
)

// defineRoutes -  Сопоставляет маршруты функцмям контроллера
func defineRoutes(router *mux.Router) {
	for _, route := range c.Routes {
		controllerFunc := c.GetFunctionByName(route.Controller)
		router.HandleFunc(route.Path, controllerFunc).Methods(route.Methods...)
	}
}

// Serve определяет пути, присоединяет функции middleware
// и запускает сервер на заданном порту.
func Serve(port string) {
	router := mux.NewRouter().StrictSlash(true)
	defineRoutes(router)
	router.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))
	router.Use(middleware.HeadersMiddleware)
	log.Fatal(http.ListenAndServe(port, router))
}
