package controller

import (
	"onlinebc_admin/model/db"

	"github.com/gin-gonic/gin"
	gq "github.com/graphql-go/graphql"
)

// TYPES ****************************************************
var broadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "Broadcast",
	Description: "Онлайн трансляция",
	Fields: gq.Fields{
		"id":             &gq.Field{Type: gq.Int, Description: "Идентификатор трансляции"},
		"title":          &gq.Field{Type: gq.String, Description: "Название трансляции"},
		"time_created":   &gq.Field{Type: gq.Int, Description: "Время создания"},
		"time_begin":     &gq.Field{Type: gq.Int, Description: "Время начала"},
		"is_ended":       &gq.Field{Type: gq.Int, Description: "Завершена 0 1"},
		"show_date":      &gq.Field{Type: gq.Int, Description: "Показывать дату 0 1"},
		"show_time":      &gq.Field{Type: gq.Int, Description: "Показывать время 0 1"},
		"is_yandex":      &gq.Field{Type: gq.Int, Description: "Яндекс трансляция 0 1"},
		"yandex_ids":     &gq.Field{Type: gq.String, Description: "Идентификаторы Яндекс трансляций"},
		"show_main_page": &gq.Field{Type: gq.Int, Description: "Показывать на главной странице 01"},
		"link_article":   &gq.Field{Type: gq.String, Description: "Ссылка на статью"},
		"link_img":       &gq.Field{Type: gq.String, Description: "Ссылка на изображение"},
		"groups_create":  &gq.Field{Type: gq.Int, Description: ""},
		"is_diary":       &gq.Field{Type: gq.Int, Description: "Дневник 01"},
		"diary_author":   &gq.Field{Type: gq.String, Description: "Автордневника"},
	},
})

var postType = gq.NewObject(gq.ObjectConfig{
	Name:        "Post",
	Description: "Пост трансляции",
	Fields: gq.Fields{
		"id":           &gq.Field{Type: gq.Int, Description: "Идентификатор поста"},
		"id_parent":    &gq.Field{Type: gq.Int, Description: "Идентификатор поста если это ответ на другой пост"},
		"id_broadcast": &gq.Field{Type: gq.Int, Description: "Идентификатор трансляции"},
		"text":         &gq.Field{Type: gq.String, Description: "Текст поста"},
		"post_time":    &gq.Field{Type: gq.Int, Description: "Текст поста"},
		"post_type":    &gq.Field{Type: gq.Int, Description: "Тип поста 1,2,3,4..."},
		"link":         &gq.Field{Type: gq.String, Description: "Ссылка"},
		"has_big_img":  &gq.Field{Type: gq.Int, Description: "Есть ли большое изображение 0,1"},
		"author":       &gq.Field{Type: gq.String, Description: "ФИО автора поста"},
	},
})

var mediumType = gq.NewObject(gq.ObjectConfig{
	Name:        "Medium",
	Description: "Медиа поста трансляции",
	Fields: gq.Fields{
		"id":      &gq.Field{Type: gq.Int, Description: "Идентификатор медиа"},
		"post_id": &gq.Field{Type: gq.Int, Description: "Идентификатор поста"},
		"uri":     &gq.Field{Type: gq.String, Description: "URI изображения"},
		"thumb":   &gq.Field{Type: gq.String, Description: "Уменьшенное изображение"},
		"source":  &gq.Field{Type: gq.String, Description: "Источник медиа"},
	},
})

// ************************************************************************

var rootQuery = gq.NewObject(gq.ObjectConfig{
	Name: "Query",
	Fields: gq.Fields{

		"broadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Показать трансляцию по идентификатору",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("broadcast", params.Args["id"].(int)), nil
			},
		},

		"post": &gq.Field{
			Type:        postType,
			Description: "Показать пост по идентификатору",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("post", params.Args["id"].(int)), nil
			},
		},

		"medium": &gq.Field{
			Type:        mediumType,
			Description: "Показать медиа по идентификатору",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("medium", params.Args["id"].(int)), nil
			},
		},

		"posts": &gq.Field{
			Type: gq.NewList(postType),
			Args: gq.FieldConfigArgument{
				"id_broadcast": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				var posts interface{}
				return posts, nil
			},
		},
	},
})

var rootMutation = gq.NewObject(gq.ObjectConfig{
	Name: "Mutation",
	Fields: gq.Fields{

		// BROADCAST =====================================================

		"createBroadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Создать трансляцию",
			Args: gq.FieldConfigArgument{
				// "id":             &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор трансляции"},
				"title":          &gq.ArgumentConfig{Type: gq.String, Description: "Название трансляции"},
				"time_created":   &gq.ArgumentConfig{Type: gq.Int, Description: "Время создания"},
				"time_begin":     &gq.ArgumentConfig{Type: gq.Int, Description: "Время начала"},
				"is_ended":       &gq.ArgumentConfig{Type: gq.Int, Description: "Завершена 0 1"},
				"show_date":      &gq.ArgumentConfig{Type: gq.Int, Description: "Показывать дату 0 1"},
				"show_time":      &gq.ArgumentConfig{Type: gq.Int, Description: "Показывать время 0 1"},
				"is_yandex":      &gq.ArgumentConfig{Type: gq.Int, Description: "Яндекс трансляция 0 1"},
				"yandex_ids":     &gq.ArgumentConfig{Type: gq.String, Description: "Идентификаторы Яндекс трансляций"},
				"show_main_page": &gq.ArgumentConfig{Type: gq.Int, Description: "Показывать на главной странице 0 1"},
				"link_article":   &gq.ArgumentConfig{Type: gq.String, Description: "Ссылка на статью"},
				"link_img":       &gq.ArgumentConfig{Type: gq.String, Description: "Ссылка на изображение"},
				"groups_create":  &gq.ArgumentConfig{Type: gq.Int, Description: ""},
				"is_diary":       &gq.ArgumentConfig{Type: gq.Int, Description: "Дневник 0 1"},
				"diary_author":   &gq.ArgumentConfig{Type: gq.String, Description: "Автор дневника"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				newRow := db.CreateRow("broadcast", params.Args)
				return newRow, nil
			},
		},

		"updateBroadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Обновить трансляцию",
			Args: gq.FieldConfigArgument{
				"id":             &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор трансляции"},
				"title":          &gq.ArgumentConfig{Type: gq.String, Description: "Название трансляции"},
				"time_created":   &gq.ArgumentConfig{Type: gq.Int, Description: "Время создания"},
				"time_begin":     &gq.ArgumentConfig{Type: gq.Int, Description: "Время начала"},
				"is_ended":       &gq.ArgumentConfig{Type: gq.Int, Description: "Завершена 0 1"},
				"show_date":      &gq.ArgumentConfig{Type: gq.Int, Description: "Показывать дату 0 1"},
				"show_time":      &gq.ArgumentConfig{Type: gq.Int, Description: "Показывать время 0 1"},
				"is_yandex":      &gq.ArgumentConfig{Type: gq.Int, Description: "Яндекс трансляция 0 1"},
				"yandex_ids":     &gq.ArgumentConfig{Type: gq.String, Description: "Идентификаторы Яндекс трансляций"},
				"show_main_page": &gq.ArgumentConfig{Type: gq.Int, Description: "Показывать на главной странице 0 1"},
				"link_article":   &gq.ArgumentConfig{Type: gq.String, Description: "Ссылка на статью"},
				"link_img":       &gq.ArgumentConfig{Type: gq.String, Description: "Ссылка на изображение"},
				"groups_create":  &gq.ArgumentConfig{Type: gq.Int, Description: ""},
				"is_diary":       &gq.ArgumentConfig{Type: gq.Int, Description: "Дневник 0 1"},
				"diary_author":   &gq.ArgumentConfig{Type: gq.String, Description: "Автор дневника"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				row := db.UpdateRowByID("broadcast", params.Args["id"].(int), params.Args)
				return row, nil
			},
		},

		"deleteBroadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Удалить трасляцию",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор трансляции"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				row := db.DeleteRowByID("broadcast", params.Args["id"].(int))
				return row, nil
			},
		},

		// POST =====================================================

		"createPost": &gq.Field{
			Type:        broadcastType,
			Description: "Создать пост к тансляции или ответ к посту",
			Args: gq.FieldConfigArgument{
				// "id":           &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор поста"}                              ,
				"id_parent":    &gq.ArgumentConfig{Type: gq.Int, Description: "Идентификатор поста если это ответ на другой пост"},
				"id_broadcast": &gq.ArgumentConfig{Type: gq.Int, Description: "Идентификатор трансляции"},
				"text":         &gq.ArgumentConfig{Type: gq.String, Description: "Текст поста"},
				"post_time":    &gq.ArgumentConfig{Type: gq.Int, Description: "Текст поста"},
				"post_type":    &gq.ArgumentConfig{Type: gq.Int, Description: "Тип поста 1,2,3,4..."},
				"link":         &gq.ArgumentConfig{Type: gq.String, Description: "Ссылка"},
				"has_big_img":  &gq.ArgumentConfig{Type: gq.Int, Description: "Есть ли большое изображение 0,1"},
				"author":       &gq.ArgumentConfig{Type: gq.String, Description: "ФИО автора поста"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				newRow := db.CreateRow("post", params.Args)
				return newRow, nil
			},
		},

		"updatePost": &gq.Field{
			Type:        broadcastType,
			Description: "Обновить пост или ответ к посту",
			Args: gq.FieldConfigArgument{
				"id":           &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор поста"},
				"id_parent":    &gq.ArgumentConfig{Type: gq.Int, Description: "Идентификатор поста если это ответ на другой пост"},
				"id_broadcast": &gq.ArgumentConfig{Type: gq.Int, Description: "Идентификатор трансляции"},
				"text":         &gq.ArgumentConfig{Type: gq.String, Description: "Текст поста"},
				"post_time":    &gq.ArgumentConfig{Type: gq.Int, Description: "Текст поста"},
				"post_type":    &gq.ArgumentConfig{Type: gq.Int, Description: "Тип поста 1,2 ,3,4..."},
				"link":         &gq.ArgumentConfig{Type: gq.String, Description: "Ссылка"},
				"has_big_img":  &gq.ArgumentConfig{Type: gq.Int, Description: "Есть ли большое изображение 0,1"},
				"author":       &gq.ArgumentConfig{Type: gq.String, Description: "ФИО автора поста"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				row := db.UpdateRowByID("post", params.Args["id"].(int), params.Args)
				return row, nil
			},
		},

		"deletePost": &gq.Field{
			Type:        broadcastType,
			Description: "Удалить пост",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор поста"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				row := db.DeleteRowByID("post", params.Args["id"].(int))
				return row, nil
			},
		},

		// MEDIA =====================================================

		"createMedium": &gq.Field{
			Type:        broadcastType,
			Description: "Создать медиа",
			Args: gq.FieldConfigArgument{
				// "id":           &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор медиа"},
				"post_id": &gq.ArgumentConfig{Type: gq.Int, Description: "Идентификатор поста"},
				"uri":     &gq.ArgumentConfig{Type: gq.String, Description: "URI изображения"},
				"thumb":   &gq.ArgumentConfig{Type: gq.String, Description: "Уменьшенное изображение"},
				"source":  &gq.ArgumentConfig{Type: gq.String, Description: "Источник медиа"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				newRow := db.CreateRow("medium", params.Args)
				return newRow, nil
			},
		},

		"updateMedium": &gq.Field{
			Type:        broadcastType,
			Description: "Обновить медиа по идентификатору",
			Args: gq.FieldConfigArgument{
				"id":      &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор медиа"},
				"post_id": &gq.ArgumentConfig{Type: gq.Int, Description: "Идентификатор поста"},
				"uri":     &gq.ArgumentConfig{Type: gq.String, Description: "URI изображения"},
				"thumb":   &gq.ArgumentConfig{Type: gq.String, Description: "Уменьшенное изображение"},
				"source":  &gq.ArgumentConfig{Type: gq.String, Description: "Источник медиа"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				row := db.UpdateRowByID("medium", params.Args["id"].(int), params.Args)
				return row, nil
			},
		},

		"deleteMedium": &gq.Field{
			Type:        broadcastType,
			Description: "Удалить медиа по идентификатору",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор медиа"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				row := db.DeleteRowByID("medium", params.Args["id"].(int))
				return row, nil
			},
		},
	},
})

var schema, _ = gq.NewSchema(gq.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

// *******************************************************************************8

// GraphQL исполняет GraphQL запрос
func (dummy) GraphQL(c *gin.Context) {
	m := getPayload(c.Request)
	result := gq.Do(gq.Params{
		Schema:        schema,
		RequestString: m["query"].(string),
	})
	c.JSON(200, result)
}
