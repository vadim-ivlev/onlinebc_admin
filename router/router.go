package router

import (
	"log"
	"net/http"
	c "onlinebc/controller"
	"onlinebc/middleware"

	"github.com/gorilla/mux"
)

// InitRoutesArray инициализирует массив маршрутов.
func InitRoutesArray() {
	c.Routes = []c.Route{
		{"Стартовая страница", "/", "/", c.LandingPage, nil},
		{"JSON  маршрутов.  Документация API.", "/routes", "/routes", c.GetRoutes, nil},
		{"Список трансляций", "/broadcasts", "/broadcasts", c.GetBroadcasts, nil},
		{"Трасляция с идентификатором id и ее постами", "/broadcast/{id}", "/broadcast/354", c.GetBroadcast, nil},
		{"Посты трансляции с идентификатором id", "/posts/{id}", "/posts/354", c.GetPosts, nil},
		{"Ответы к посту с идентификатором id", "/answers/{id}", "/answers/23932", c.GetAnswers, nil},
		{"Медиа поста с идентификатором id", "/media/{id}", "/media/23932", c.GetMedia, nil},
		{"Трасляция с идентификатором id и ее постами. Legacy", "/api/online.php", "/api/online.php?id=354", c.GetBroadcast, []c.Param{
			{"Идентификатор трансляции", "id", "{id}"},
		}},
		{"Список трансляций.Legacy", "/api/", "/api/?main=0&active=0&num=3", c.GetBroadcastList, []c.Param{
			{"Основная {0|1}", "main", "{main}"},
			{"Активность {0|1}", "active", "{active}"},
			{"Номер", "num", "{num}"},
		}},
	}
}

// defineRoutes -  Сопоставляет маршруты контроллерам для заданного раутера
func defineRoutes(router *mux.Router) {

	for _, route := range c.Routes {
		r := router.HandleFunc(route.Path, route.Func).Methods("GET", "HEAD")
		for _, param := range route.Params {
			r.Queries(param.Name, param.Value)
		}
	}
}

// Serve определяет пути, присоединяет функции middleware
// и запускает сервер на заданном порту.
func Serve(port string) {
	router := mux.NewRouter()
	InitRoutesArray()
	defineRoutes(router)
	router.Use(middleware.HeadersMiddleware)
	router.Use(middleware.RedisMiddleware)
	log.Fatal(http.ListenAndServe(port, router))
}
