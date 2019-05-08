package controller

import (
	"encoding/json"

	gq "github.com/graphql-go/graphql"
)

// F U N C S ***********************************************

// JSONParamToMap - возвращает параметр paramName в map[string]interface{}.
// Второй параметр возврата - ошибка.
// Применяется для сериализации поля JSON таблицы postgres в map.
func JSONParamToMap(params gq.ResolveParams, paramName string) (interface{}, error) {
	var paramMap []map[string]interface{}
	source := params.Source.(map[string]interface{})
	paramBytes, ok := source[paramName].([]byte)
	if !ok {
		return paramMap, nil
	}
	err := json.Unmarshal(paramBytes, &paramMap)
	return paramMap, err
}

type fields map[string]*gq.Field

func (sumFields fields) addFields(fields1 fields) {
	for key, field := range fields1 {
		sumFields[key] = field
	}
}

func addFields(fields1 map[string]*gq.Field, fields2 map[string]*gq.Field) map[string]*gq.Field {
	sumFields := make(map[string]*gq.Field)

	// for key, field := range fields1 {
	// 	sumFields[key] = field
	// }
	// for key, field := range fields2 {
	// 	sumFields[key] = field
	// }

	fields(sumFields).addFields(fields(fields1))
	fields(sumFields).addFields(fields(fields2))

	return sumFields
}

// FIELDS **************************************************
var broadcastFields = gq.Fields{
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
}

var postFields = gq.Fields{
	"id": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор поста",
	},
	"id_parent": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор родительского поста если это ответ на другой пост",
	},
	"id_broadcast": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор трансляции",
	},
	"text": &gq.Field{
		Type:        gq.String,
		Description: "Текст поста",
	},
	"post_time": &gq.Field{
		Type:        gq.Int,
		Description: "Текст поста",
	},
	"post_type": &gq.Field{
		Type:        gq.Int,
		Description: "Тип поста 1,2,3,4...",
	},
	"link": &gq.Field{
		Type:        gq.String,
		Description: "Ссылка",
	},
	"has_big_img": &gq.Field{
		Type:        gq.Int,
		Description: "Есть ли большое изображение 0,1",
	},
	"author": &gq.Field{
		Type:        gq.String,
		Description: "ФИО автора поста",
	},
}

var mediumFields = gq.Fields{
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
}

var listBroadcastFields = gq.Fields{
	"length": &gq.Field{
		Type:        gq.Int,
		Description: "Количество элементов в списке",
	},
	"list": &gq.Field{
		Type:        gq.NewList(broadcastType),
		Description: "Список трансляций",
	},
}

// FULL FIELDS поля с древовидной структурой  ****************************************************

// var fullAnswerFields = addFields(postFields, gq.Fields{
// 	"media": &gq.Field{
// 		Type:        gq.NewList(mediumType),
// 		Description: "Медиа ответа",
// 	},
// })

var fullAnswerFields = gq.Fields{
	"id": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор ответа",
	},
	"id_parent": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор родительского поста если это ответ на другой пост",
	},
	"id_broadcast": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор трансляции",
	},
	"text": &gq.Field{
		Type:        gq.String,
		Description: "Текст поста",
	},
	"post_time": &gq.Field{
		Type:        gq.Int,
		Description: "Текст поста",
	},
	"post_type": &gq.Field{
		Type:        gq.Int,
		Description: "Тип поста 1,2,3,4...",
	},
	"link": &gq.Field{
		Type:        gq.String,
		Description: "Ссылка",
	},
	"has_big_img": &gq.Field{
		Type:        gq.Int,
		Description: "Есть ли большое изображение 0,1",
	},
	"author": &gq.Field{
		Type:        gq.String,
		Description: "ФИО автора поста",
	},
	// ---------------------------------
	"media": &gq.Field{
		Type:        gq.NewList(mediumType),
		Description: "Медиа ответа",
		Resolve: func(params gq.ResolveParams) (interface{}, error) {
			return JSONParamToMap(params, "media")
		},
	},
}

// var fullPostFields = addFields(postFields, gq.Fields{
// 	"media": &gq.Field{
// 		Type:        gq.NewList(mediumType),
// 		Description: "Медиа поста",
// 	},
// 	"answers": &gq.Field{
// 		Type:        gq.NewList(fullAnswerType),
// 		Description: "Ответы к посту",
// 	},
// })

var fullPostFields = gq.Fields{
	"id": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор поста",
	},
	"id_parent": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор поста если это ответ на другой пост",
	},
	"id_broadcast": &gq.Field{
		Type:        gq.Int,
		Description: "Идентификатор трансляции",
	},
	"text": &gq.Field{
		Type:        gq.String,
		Description: "Текст поста",
	},
	"post_time": &gq.Field{
		Type:        gq.Int,
		Description: "Текст поста",
	},
	"post_type": &gq.Field{
		Type:        gq.Int,
		Description: "Тип поста 1,2,3,4...",
	},
	"link": &gq.Field{
		Type:        gq.String,
		Description: "Ссылка",
	},
	"has_big_img": &gq.Field{
		Type:        gq.Int,
		Description: "Есть ли большое изображение 0,1",
	},
	"author": &gq.Field{
		Type:        gq.String,
		Description: "ФИО автора поста",
	},
	// --------------------------------------------
	"media": &gq.Field{
		Type:        gq.NewList(mediumType),
		Description: "Медиа поста",
		Resolve: func(params gq.ResolveParams) (interface{}, error) {
			return JSONParamToMap(params, "media")
		},
	},
	"answers": &gq.Field{
		Type:        gq.NewList(fullAnswerType),
		Description: "Ответы к посту",
		Resolve: func(params gq.ResolveParams) (interface{}, error) {
			return JSONParamToMap(params, "answers")
		},
	},
}

// var fullBroadcastFields = addFields(broadcastFields, gq.Fields{
// 	"posts": &gq.Field{
// 		Type:        gq.NewList(fullPostType),
// 		Description: "Посты бродкаста",
// 	},
// })

var fullBroadcastFields = gq.Fields{
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
	// ------------------------------------
	"posts": &gq.Field{
		Type:        gq.NewList(fullPostType),
		Description: "Посты бродкаста",
		Resolve: func(params gq.ResolveParams) (interface{}, error) {
			return JSONParamToMap(params, "posts")
		},
	},
}

var FullListBroadcastFields = gq.Fields{
	"length": &gq.Field{
		Type:        gq.Int,
		Description: "Количество элементов в списке",
	},
	"list": &gq.Field{
		Type:        gq.NewList(fullBroadcastType),
		Description: "Список трансляций c постами, ответами и медиа",
	},
}

// TYPES ****************************************************

var postType = gq.NewObject(gq.ObjectConfig{
	Name:        "Post",
	Description: "Пост трансляции",
	Fields:      postFields,
})

var mediumType = gq.NewObject(gq.ObjectConfig{
	Name:        "Medium",
	Description: "Медиа поста трансляции",
	Fields:      mediumFields,
})

var broadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "Broadcast",
	Description: "Онлайн трансляция",
	Fields:      broadcastFields,
})

var listBroadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "ListBroadcast",
	Description: "Список трансляций и количество элементов в списке",
	Fields:      listBroadcastFields,
})

// FULL TYPES типы с древовидной структурой *************

var fullPostType = gq.NewObject(gq.ObjectConfig{
	Name:        "FullPost",
	Description: "Пост трансляции с медиа и ответами к посту",
	Fields:      fullPostFields,
})

var fullAnswerType = gq.NewObject(gq.ObjectConfig{
	Name:        "FullAnswer",
	Description: "Ответ к посту с медиа ответа",
	Fields:      fullAnswerFields,
})

var fullBroadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "FullBroadcast",
	Description: "Трансляция с постами",
	Fields:      fullBroadcastFields,
})

var fullListBroadcastType = gq.NewObject(gq.ObjectConfig{
	Name:        "FullListBroadcast",
	Description: "Список трансляций c постами, ответами и медиа,  и количество элементов в списке",
	Fields:      FullListBroadcastFields,
})
