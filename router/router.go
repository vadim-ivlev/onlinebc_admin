package router

import (
	c "onlinebc_admin/controller"
	"onlinebc_admin/middleware"

	"github.com/gin-gonic/gin"
)

// defineRoutes -  Сопоставляет маршруты функцмям контроллера
func defineGinRoutes(router *gin.Engine) {
	for _, route := range c.Routes {
		controllerFunc := c.GetGinFunctionByName(route.Controller)
		for _, method := range route.Methods {
			router.Handle(method, route.Path, controllerFunc)
		}
	}
}

// GinServe определяет пути, присоединяет функции middleware
// и запускает сервер на заданном порту.
func GinServe(port string) {
	router := gin.Default()
	router.Use(middleware.GinHeadersMiddleware())

	router.LoadHTMLGlob("templates/*.*")

	defineGinRoutes(router)
	// router.Static("/templates", "./templates")
	router.Run(port)
}
