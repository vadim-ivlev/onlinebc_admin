package router

import (
	c "onlinebc_admin/controller"
	"onlinebc_admin/middleware"

	"github.com/gin-gonic/gin"
)

// routeExists проверяет отсутствует ли в раутере пара метод+путь
func routeAbsent(router *gin.Engine, method string, path string) bool {
	routes := router.Routes()
	for _, r := range routes {
		if r.Method == method && r.Path == path {
			return false
		}
	}
	return true
}

// defineRoutes -  Сопоставляет маршруты функцмям контроллера
func defineRoutes(router *gin.Engine) {
	for _, route := range c.Routes {
		controllerFunc := c.GetFunctionByName(route.Controller)
		for _, method := range route.Methods {
			if routeAbsent(router, method, route.Path) {
				router.Handle(method, route.Path, controllerFunc)
			}

		}
	}
}

// Serve определяет пути, присоединяет функции middleware
// и запускает сервер на заданном порту.
func Serve(port string, debug bool) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.StaticFile("/favicon.ico", "./templates/favicon.ico")
	router.LoadHTMLGlob("templates/*.*")

	//Middleware
	router.Use(middleware.HeadersMiddleware())
	router.Use(middleware.RedisMiddleware())

	defineRoutes(router)
	router.Run(port)
}
