package router

import (
	"log"
	"net/http"
	c "onlinebc_admin/controller"
	"onlinebc_admin/middleware"

	"github.com/gorilla/mux"
)

// InitRoutesArray инициализирует массив маршрутов.
func InitRoutesArray() {
	c.Routes = []c.Route{

		{[]string{"GET"}, "Получить Медиа с идентификатором id", "/medium/{id}", "/medium/5330", c.GetMedium, nil},
		{[]string{"GET"}, "Получить Пост с идентификатором id", "/post/{id}", "/post/23932", c.GetPost, nil},
		{[]string{"GET"}, "Получить Трансляцию с идентификатором id", "/broadcast/{id}", "/broadcast/354", c.GetBroadcast, nil},

		{[]string{"PUT"}, "Редактировать Медиа с идентификатором id", "/update-medium/{id}", "/update-medium/5330?uri=uriimg&thumb=t&source=somesource", c.UpdateMedium, []c.Param{
			{"URI изображения", "uri", "{uri}"},
			{"Уменьшенное изображение", "thumb", "{thumb}"},
			{"Источник", "source", "{source}"},
		}},

		{[]string{"PUT"}, "Редактировать Пост с идентификатором id", "/post/{id}", "/post/23932", c.GetPost, nil},
		{[]string{"PUT"}, "Редактировать Трансляцию с идентификатором id", "/broadcast/{id}", "/broadcast/354", c.GetBroadcast, nil},

		{[]string{"POST"}, "Создать Медиа с идентификатором id", "/medium/{id}", "/medium/5330", c.GetMedium, nil},
		{[]string{"POST"}, "Создать Пост с идентификатором id", "/post/{id}", "/post/23932", c.GetPost, nil},
		{[]string{"POST"}, "Создать Трансляцию с идентификатором id", "/broadcast/{id}", "/broadcast/354", c.GetBroadcast, nil},

		{[]string{"DELETE"}, "Удалить Медиа с идентификатором id", "/medium/{id}", "/medium/5330", c.GetMedium, nil},
		{[]string{"DELETE"}, "Удалить Пост с идентификатором id", "/post/{id}", "/post/23932", c.GetPost, nil},
		{[]string{"DELETE"}, "Удалить Трансляцию с идентификатором id", "/broadcast/{id}", "/broadcast/354", c.GetBroadcast, nil},

		{[]string{"GET", "HEAD"}, "Стартовая страница", "/", "/", c.LandingPage, nil},
		{[]string{"GET", "HEAD"}, "JSON  маршрутов.  Документация API.", "/routes", "/routes", c.GetRoutes, nil},
		{[]string{"GET", "HEAD"}, "Список трансляций", "/broadcasts", "/broadcasts", c.GetBroadcasts, nil},
		{[]string{"GET", "HEAD"}, "Трасляция с идентификатором id и ее постами", "/full-broadcast/{id}", "/full-broadcast/354", c.GetFullBroadcast, nil},
		{[]string{"GET", "HEAD"}, "Посты трансляции с идентификатором id", "/posts/{id}", "/posts/354", c.GetPosts, nil},
		{[]string{"GET", "HEAD"}, "Ответы к посту с идентификатором id", "/answers/{id}", "/answers/23932", c.GetAnswers, nil},
		{[]string{"GET", "HEAD"}, "Медиа записи поста с идентификатором id", "/media/{id}", "/media/23932", c.GetMedia, nil},
		{[]string{"GET", "HEAD"}, "Трасляция с идентификатором id и ее постами. Legacy", "/api/online.php", "/api/online.php?id=354", c.GetFullBroadcast, []c.Param{
			{"Идентификатор трансляции", "id", "{id}"},
		}},
		{[]string{"GET", "HEAD"}, "Список трансляций.Legacy", "/api/", "/api/?main=0&active=0&num=3", c.GetBroadcastList, []c.Param{
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
	log.Fatal(http.ListenAndServe(port, router))
}
