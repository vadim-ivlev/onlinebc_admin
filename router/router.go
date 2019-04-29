package router

import (
	"net/http"
	c "onlinebc_admin/controller"
	"onlinebc_admin/middleware"

	"github.com/gin-gonic/gin"
)

// routeAbsent проверяет отсутствует ли в раутере пара метод+путь
func routeAbsent(r *gin.Engine, method string, path string) bool {
	routes := r.Routes()
	for _, r := range routes {
		if r.Method == method && r.Path == path {
			return false
		}
	}
	return true
}

// defineRoutes -  Сопоставляет маршруты функцмям контроллера
func defineRoutes(r *gin.Engine) {
	r.Handle("OPTIONS", "/graphql", PingHandler)

	for _, route := range c.Routes {
		controllerFunc := c.GetFunctionByName(route.Controller)
		for _, method := range route.Methods {
			if routeAbsent(r, method, route.Path) {
				r.Handle(method, route.Path, controllerFunc)
			}

		}
	}
}

// PingHandler нужен для фронта, так как сначала отправляется метод с OPTIONS
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// Setup определяет пути и присоединяет функции middleware.
// Если debug == true в консоль выдается больше информации.
// Если outputToConsole == false вывод в консоль не производится
func Setup(debug bool, outputToConsole bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var r *gin.Engine
	if outputToConsole {
		r = gin.Default()
	} else {
		r = gin.New()
	}

	r.StaticFile("/favicon.ico", "./templates/favicon.ico")
	// r.Static("/uploads_temp", "./uploads_temp")
	r.LoadHTMLGlob("templates/*.*")

	// подключаем Middleware
	r.Use(middleware.HeadersMiddleware())
	r.Use(middleware.RedisMiddleware())

	defineRoutes(r)
	return r
}

// Serve запускает сервер на заданном порту.
// func Serve(port string, debug bool) {
// 	r := SetupRouter(debug, true)
// 	r.Run(port)
// }
