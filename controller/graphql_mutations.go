package controller

import (
	"log"
	"onlinebc_admin/model/db"

	gq "github.com/graphql-go/graphql"
)

var rootMutation = gq.NewObject(gq.ObjectConfig{
	Name: "Mutation",
	Fields: gq.Fields{

		// BROADCAST =====================================================

		"create_broadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Создать трансляцию",
			Args: gq.FieldConfigArgument{
				// "id":             &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор трансляции"},
				"title": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Название трансляции",
				},
				"time_created": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Время создания",
				},
				"time_begin": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Время начала",
				},
				"is_ended": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "Завершена 0 1. По умолчанию 0",
					DefaultValue: 0,
				},
				"show_date": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Показывать дату 0 1",
				},
				"show_time": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Показывать время 0 1",
				},
				"show_main_page": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Показывать на главной странице 0 1",
				},
				"link_article": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Ссылка на статью",
				},
				"link_img": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Ссылка на изображение",
				},
				"groups_create": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "",
				},
				"is_diary": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Дневник 0 1",
				},
				"diary_author": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Автор дневника",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.CreateRow("broadcast", params.Args)
			},
		},

		"update_broadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Обновить трансляцию",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор трансляции",
				},
				"title": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Название трансляции",
				},
				"time_created": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Время создания",
				},
				"time_begin": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Время начала",
				},
				"is_ended": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Завершена 0 1",
				},
				"show_date": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Показывать дату 0 1",
				},
				"show_time": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Показывать время 0 1",
				},
				"show_main_page": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Показывать на главной странице 0 1",
				},
				"link_article": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Ссылка на статью",
				},
				"link_img": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Ссылка на изображение",
				},
				"groups_create": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "",
				},
				"is_diary": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Дневник 0 1",
				},
				"diary_author": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Автор дневника",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.UpdateRowByID("broadcast", params.Args["id"].(int), params.Args)
			},
		},

		"delete_broadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Удалить трансляцию",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор трансляции",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.DeleteRowByID("broadcast", params.Args["id"].(int))
			},
		},

		// POST =====================================================

		"create_post": &gq.Field{
			Type:        postType,
			Description: "Создать пост к трансляции или ответ к посту",
			Args: gq.FieldConfigArgument{
				// "id":           &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор поста"}                              ,
				"id_parent": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор поста если это ответ на другой пост",
				},
				"id_broadcast": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор трансляции",
				},
				"text": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Текст поста",
				},
				"post_time": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Текст поста",
				},
				"post_type": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Тип поста 1,2,3,4...",
				},
				"link": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Ссылка",
				},
				"has_big_img": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Есть ли большое изображение 0,1",
				},
				"author": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "ФИО автора поста",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.CreateRow("post", params.Args)
			},
		},

		"update_post": &gq.Field{
			Type:        postType,
			Description: "Обновить пост или ответ к посту",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор поста",
				},
				"id_parent": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор поста если это ответ на другой пост",
				},
				"id_broadcast": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор трансляции",
				},
				"text": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Текст поста",
				},
				"post_time": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Текст поста",
				},
				"post_type": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Тип поста 1,2 ,3,4...",
				},
				"link": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Ссылка",
				},
				"has_big_img": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Есть ли большое изображение 0,1",
				},
				"author": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "ФИО автора поста",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.UpdateRowByID("post", params.Args["id"].(int), params.Args)
			},
		},

		"delete_post": &gq.Field{
			Type:        postType,
			Description: "Удалить пост",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор поста"},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.DeleteRowByID("post", params.Args["id"].(int))
			},
		},

		// MEDIA =====================================================

		"create_image": &gq.Field{
			Type:        imageType,
			Description: "Создать медиа",
			Args: gq.FieldConfigArgument{
				// "id":           &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int), Description: "Идентификатор изображения"},
				"post_id": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор поста",
				},
				"filepath": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "URI изображения",
				},
				"source": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Источник медиа",
				},
				"file_field_name": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Имя (name) поля формы для загрузки файла. <input name='fname' type='file' ...>",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {

				path, _, _, thumbs, errMsg := SaveUploadedImage(params, "file_field_name")
				if errMsg == "" {
					params.Args["filepath"] = path
					// params.Args["width"] = width
					// params.Args["height"] = height
					params.Args["thumbs"] = thumbs
				} else {
					log.Println("create_image: Resolve(): " + errMsg)
				}
				delete(params.Args, "file_field_name")
				return db.CreateRow("image", params.Args)

			},
		},

		"update_image": &gq.Field{
			Type:        imageType,
			Description: "Обновить медиа по идентификатору",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int),
					Description: "Идентификатор медиа",
				},
				"post_id": &gq.ArgumentConfig{Type: gq.Int,
					Description: "Идентификатор поста",
				},
				"filepath": &gq.ArgumentConfig{Type: gq.String,
					Description: "URI изображения",
				},
				"source": &gq.ArgumentConfig{Type: gq.String,
					Description: "Источник медиа",
				},
				"file_field_name": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Имя (name) поля формы для загрузки файла. <input name='fname' type='file' ...>",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {

				path, _, _, thumbs, errMsg := SaveUploadedImage(params, "file_field_name")
				if errMsg == "" {
					params.Args["filepath"] = path
					// params.Args["width"] = width
					// params.Args["height"] = height
					params.Args["thumbs"] = thumbs
				} else {
					log.Println("create_image: Resolve(): " + errMsg)
				}
				delete(params.Args, "file_field_name")
				return db.UpdateRowByID("image", params.Args["id"].(int), params.Args)
			},
		},

		"delete_image": &gq.Field{
			Type:        imageType,
			Description: "Удалить медиа по идентификатору",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор медиа",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.DeleteRowByID("image", params.Args["id"].(int))
			},
		},
	},
})
