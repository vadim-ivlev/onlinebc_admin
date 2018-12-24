package router

import (
	"fmt"
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
		r := router.HandleFunc(route.Path, controllerFunc).Methods(route.Methods...)
		mm, _ := r.GetMethods()
		fmt.Printf("%v", mm)
		for _, param := range route.Params {
			r.Queries(param.Name, param.Value)
		}
	}
}

// Serve определяет пути, присоединяет функции middleware
// и запускает сервер на заданном порту.
func Serve(port string) {
	router := mux.NewRouter().StrictSlash(true)
	defineRoutes(router)
	router.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	router.Use(middleware.HeadersMiddleware)
	log.Fatal(http.ListenAndServe(port, router))
}
