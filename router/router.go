package router

import (
	"fmt"
	"log"
	"net/http"
	c "onlinebc_admin/controller"
	"onlinebc_admin/middleware"

	"github.com/gin-gonic/gin"
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

// defineRoutes -  Сопоставляет маршруты функцмям контроллера
func defineGinRoutes(router *gin.Engine) {
	for _, route := range c.Routes {
		// controllerFunc := c.GetFunctionByName(route.Controller)
		for _, method := range route.Methods {
			// router.Handle(method, route.Path, controllerFunc)
			fmt.Println(method, route.Controller)
		}
	}
}

// GinServe определяет пути, присоединяет функции middleware
// и запускает сервер на заданном порту.
func GinServe(port string) {
	router := gin.Default()
	defineGinRoutes(router)
	router.Static("/templates", "./templates")
	// router.Use(middleware.HeadersMiddleware)
	router.Run(port)
}
