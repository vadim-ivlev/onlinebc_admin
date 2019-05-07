package controller

import (
	gq "github.com/graphql-go/graphql"
)

// TYPES ****************************************************
var broadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "Broadcast",
	Description: "Онлайн трансляция",
	Fields: gq.Fields{
		"id": &gq.Field{
			Type:        gq.Int,
			Description: "Идентификатор трансляции",
		},
		"title": &gq.Field{
			Type:        gq.String,
			Description: "Название трансляции",
		},
		"time_created": &gq.Field{
			Type:        gq.Int,
			Description: "Время создания",
		},
		"time_begin": &gq.Field{
			Type:        gq.Int,
			Description: "Время начала",
		},
		"is_ended": &gq.Field{
			Type:        gq.Int,
			Description: "Завершена 0 1",
		},
		"show_date": &gq.Field{
			Type:        gq.Int,
			Description: "Показывать дату 0 1",
		},
		"show_time": &gq.Field{
			Type:        gq.Int,
			Description: "Показывать время 0 1",
		},
		"show_main_page": &gq.Field{
			Type:        gq.Int,
			Description: "Показывать на главной странице 01",
		},
		"link_article": &gq.Field{
			Type:        gq.String,
			Description: "Ссылка на статью",
		},
		"link_img": &gq.Field{
			Type:        gq.String,
			Description: "Ссылка на изображение",
		},
		"groups_create": &gq.Field{
			Type:        gq.Int,
			Description: "",
		},
		"is_diary": &gq.Field{
			Type:        gq.Int,
			Description: "Дневник 01",
		},
		"diary_author": &gq.Field{
			Type:        gq.String,
			Description: "Автордневника",
		},
	},
})

var postType = gq.NewObject(gq.ObjectConfig{
	Name:        "Post",
	Description: "Пост трансляции",
	Fields: gq.Fields{
		"id": &gq.Field{
			Type:        gq.Int,
			Description: "Идентификатор поста",
		},
		"id_parent": &gq.Field{Type: gq.Int,
			Description: "Идентификатор поста если это ответ на другой пост"},
		"id_broadcast": &gq.Field{Type: gq.Int,
			Description: "Идентификатор трансляции"},
		"text": &gq.Field{Type: gq.String,
			Description: "Текст поста"},
		"post_time": &gq.Field{Type: gq.Int,
			Description: "Текст поста"},
		"post_type": &gq.Field{Type: gq.Int,
			Description: "Тип поста 1,2,3,4..."},
		"link": &gq.Field{Type: gq.String,
			Description: "Ссылка"},
		"has_big_img": &gq.Field{Type: gq.Int,
			Description: "Есть ли большое изображение 0,1"},
		"author": &gq.Field{Type: gq.String,
			Description: "ФИО автора поста"},
	},
})

var mediumType = gq.NewObject(gq.ObjectConfig{
	Name:        "Medium",
	Description: "Медиа поста трансляции",
	Fields: gq.Fields{
		"id": &gq.Field{
			Type:        gq.Int,
			Description: "Идентификатор медиа",
		},
		"post_id": &gq.Field{
			Type:        gq.Int,
			Description: "Идентификатор поста",
		},
		"uri": &gq.Field{
			Type:        gq.String,
			Description: "URI изображения",
		},
		"thumb": &gq.Field{
			Type:        gq.String,
			Description: "Уменьшенное изображение",
		},
		"source": &gq.Field{
			Type:        gq.String,
			Description: "Источник медиа",
		},
	},
})

var listBroadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "ListBroadcast",
	Description: "Список трансляций и количество элементов в списке",
	Fields: gq.Fields{
		"length": &gq.Field{
			Type:        gq.Int,
			Description: "Количество элементов в списке",
		},
		"list": &gq.Field{
			Type:        gq.NewList(broadcastType),
			Description: "Список трансляций",
		},
	},
})
