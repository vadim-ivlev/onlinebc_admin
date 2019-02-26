package router

import (
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
	for _, route := range c.Routes {
		controllerFunc := c.GetFunctionByName(route.Controller)
		for _, method := range route.Methods {
			if routeAbsent(r, method, route.Path) {
				r.Handle(method, route.Path, controllerFunc)
			}

		}
	}
}

// SetupRouter определяет пути и присоединяет функции middleware.
// Если debug == true в консоль выдается больше информации.
// Если outputToConsole == false вывод в консоль не производится
func SetupRouter(debug bool, outputToConsole bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var r *gin.Engine
	if outputToConsole {
		r = gin.Default()
	} else {
		r = gin.New()
	}
	// TODO: increase max size of POST
	r.MaxMultipartMemory = 100 << 20 // 100 * 1024 * 1024 // 100 MiB

	r.StaticFile("/favicon.ico", "./templates/favicon.ico")
	r.Static("/uploaded_images", "./uploaded_images")
	r.LoadHTMLGlob("templates/*.*")

	// подключаем Middleware
	r.Use(middleware.HeadersMiddleware())
	r.Use(middleware.RedisMiddleware())

	defineRoutes(r)
	return r
}

// Serve запускает сервер на заданном порту.
func Serve(port string, debug bool) {
	r := SetupRouter(debug, true)
	r.Run(port)
}
