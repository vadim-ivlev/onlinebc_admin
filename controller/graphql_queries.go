package controller

import (
	"fmt"
	"onlinebc_admin/model/db"

	gq "github.com/graphql-go/graphql"
)

// ************************************************************************

var rootQuery = gq.NewObject(gq.ObjectConfig{
	Name: "Query",
	Fields: gq.Fields{

		"get_broadcast": &gq.Field{
			Type:        broadcastType,
			Description: "Показать трансляцию по идентификатору",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("broadcast", params.Args["id"].(int))
			},
		},

		"get_full_broadcast": &gq.Field{
			Type:        fullBroadcastType,
			Description: "Показать трансляцию по идентификатору c постами, ответами и медиа",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("full_broadcast", params.Args["id"].(int))
			},
		},

		"get_post": &gq.Field{
			Type:        postType,
			Description: "Показать пост по идентификатору",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("post", params.Args["id"].(int))
			},
		},

		"get_full_post": &gq.Field{
			Type:        fullPostType,
			Description: "Показать пост с ответами и изображениями по идентификатору поста",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("full_post", params.Args["id"].(int))
			},
		},

		"get_medium": &gq.Field{
			Type:        mediumType,
			Description: "Показать медиа по идентификатору",
			Args:        gq.FieldConfigArgument{"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.Int)}},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.GetRowByID("medium", params.Args["id"].(int))
			},
		},

		"get_broadcast_posts": &gq.Field{
			Type:        gq.NewList(postType),
			Description: "Получить посты трансляции по ее идентификатору.",
			Args: gq.FieldConfigArgument{
				"id_broadcast": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор трансляции",
				},
				"show_answers": &gq.ArgumentConfig{
					Type:         gq.Boolean,
					DefaultValue: false,
					Description:  "Идентификатор трансляции",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				showAnswersCondition := "AND id_parent IS NULL"
				if params.Args["show_answers"].(bool) {
					showAnswersCondition = ""
				}
				return db.QuerySliceMap("SELECT * FROM post WHERE id_broadcast = $1 "+showAnswersCondition+" ORDER BY post_time DESC ;", params.Args["id_broadcast"].(int))
			},
		},

		"get_post_media": &gq.Field{
			Type:        gq.NewList(mediumType),
			Description: "Получить фотографии поста по его идентификатору.",
			Args: gq.FieldConfigArgument{
				"post_id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор поста",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.QuerySliceMap("SELECT * FROM medium WHERE post_id = $1 ;", params.Args["post_id"].(int))
			},
		},

		"get_post_answers": &gq.Field{
			Type:        gq.NewList(postType),
			Description: "Получить ответы к посту по идентификатору поста.",
			Args: gq.FieldConfigArgument{
				"id_parent": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор поста",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return db.QuerySliceMap("SELECT * FROM post WHERE id_parent = $1 ;", params.Args["id_parent"].(int))
			},
		},

		// "broadcasts": &gq.Field{
		// 	Type:        gq.NewList(broadcastType),
		// 	Description: "Получить список трансляций.",
		// 	Args: gq.FieldConfigArgument{
		// 		"search": &gq.ArgumentConfig{
		// 			Type:         gq.String,
		// 			Description:  "Строка полнотекстового поиска. По умолчанию ''.",
		// 			DefaultValue: "",
		// 		},
		// 		"is_ended": &gq.ArgumentConfig{
		// 			Type:         gq.Int,
		// 			Description:  "1 если трансляция закончена, 0 - если нет. По умолчанию 1.",
		// 			DefaultValue: 1,
		// 		},
		// 		"order": &gq.ArgumentConfig{
		// 			Type:         gq.String,
		// 			Description:  "сортировка строк в определённом порядке. По умолчанию 'id DESC'",
		// 			DefaultValue: "id DESC",
		// 		},
		// 		"limit": &gq.ArgumentConfig{
		// 			Type:         gq.Int,
		// 			Description:  "возвратить не больше заданного числа строк. По умолчанию 100.",
		// 			DefaultValue: 100,
		// 		},
		// 		"offset": &gq.ArgumentConfig{
		// 			Type:         gq.Int,
		// 			Description:  "пропустить указанное число строк, прежде чем начать выдавать строки. По умолчанию 0.",
		// 			DefaultValue: 0,
		// 		},
		// 	},
		// 	Resolve: func(params gq.ResolveParams) (interface{}, error) {
		// 		// s := params.Args["search"].(string)
		// 		// textSearchCondition := ""
		// 		// if len(s) > 0 {
		// 		// 	textSearchCondition = fmt.Sprintf("to_tsvector('russian', title) @@ plainto_tsquery('russian','%s') AND", s)
		// 		// }

		// 		// return db.QuerySliceMap("SELECT * FROM broadcast WHERE "+textSearchCondition+
		// 		// 	" is_ended = $1 ORDER BY $2 LIMIT $3 OFFSET $4 ;",
		// 		// 	params.Args["is_ended"].(int),
		// 		// 	params.Args["order"].(string),
		// 		// 	params.Args["limit"].(int),
		// 		// 	params.Args["offset"].(int),
		// 		// )

		// 		wherePart, orderAndLimits := queryEnd(params)
		// 		return db.QuerySliceMap("SELECT * FROM broadcast" + wherePart + orderAndLimits)

		// 	},
		// },

		"list_broadcast": &gq.Field{
			Type:        listBroadcastType,
			Description: "Получить список трансляций и их количество.",
			Args: gq.FieldConfigArgument{
				"search": &gq.ArgumentConfig{
					Type:         gq.String,
					Description:  "Строка полнотекстового поиска. По умолчанию ''.",
					DefaultValue: "",
				},
				"is_ended": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "1 если трансляция закончена, 0 - если нет. По умолчанию 1.",
					DefaultValue: 1,
				},
				"order": &gq.ArgumentConfig{
					Type:         gq.String,
					Description:  "сортировка строк в определённом порядке. По умолчанию 'id DESC'",
					DefaultValue: "id DESC",
				},
				"limit": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "возвратить не больше заданного числа строк. По умолчанию 100.",
					DefaultValue: 100,
				},
				"offset": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "пропустить указанное число строк, прежде чем начать выдавать строки. По умолчанию 0.",
					DefaultValue: 0,
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				wherePart, orderAndLimits := queryEnd(params)

				list, err := db.QuerySliceMap("SELECT * FROM broadcast" + wherePart + orderAndLimits)
				if err != nil {
					return nil, err
				}
				count, err := db.QueryRowMap("SELECT count(*) AS count FROM broadcast" + wherePart)
				if err != nil {
					return nil, err
				}

				length := count["count"]

				m := map[string]interface{}{
					"length": length,
					"list":   list,
				}

				return m, nil

			},
		},

		"list_full_broadcast": &gq.Field{
			Type:        fullListBroadcastType,
			Description: "Получить список трансляций c постами, ответами и изображениями, и их количество.",
			Args: gq.FieldConfigArgument{
				"search": &gq.ArgumentConfig{
					Type:         gq.String,
					Description:  "Строка полнотекстового поиска. По умолчанию ''.",
					DefaultValue: "",
				},
				"is_ended": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "1 если трансляция закончена, 0 - если нет. По умолчанию 1.",
					DefaultValue: 1,
				},
				"order": &gq.ArgumentConfig{
					Type:         gq.String,
					Description:  "сортировка строк в определённом порядке. По умолчанию 'id DESC'",
					DefaultValue: "id DESC",
				},
				"limit": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "возвратить не больше заданного числа строк. По умолчанию 100.",
					DefaultValue: 100,
				},
				"offset": &gq.ArgumentConfig{
					Type:         gq.Int,
					Description:  "пропустить указанное число строк, прежде чем начать выдавать строки. По умолчанию 0.",
					DefaultValue: 0,
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				wherePart, orderAndLimits := queryEnd(params)

				list, err := db.QuerySliceMap("SELECT * FROM full_broadcast" + wherePart + orderAndLimits)
				if err != nil {
					return nil, err
				}
				count, err := db.QueryRowMap("SELECT count(*) AS count FROM full_broadcast" + wherePart)
				if err != nil {
					return nil, err
				}

				length := count["count"]

				m := map[string]interface{}{
					"length": length,
					"list":   list,
				}

				return m, nil

			},
		},
	},
})

// queryEnd возвращает вторую часть запроса на поиск трансляций
func queryEnd(params gq.ResolveParams) (wherePart string, orderAndLimits string) {
	s := params.Args["search"].(string)
	textSearchCondition := ""
	if len(s) > 0 {
		textSearchCondition = fmt.Sprintf("to_tsvector('russian', title) @@ plainto_tsquery('russian','%s') AND", s)
	}

	wherePart = fmt.Sprintf(" WHERE %s is_ended = %d ",
		textSearchCondition,
		params.Args["is_ended"].(int),
	)
	orderAndLimits = fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d ;",
		params.Args["order"].(string),
		params.Args["limit"].(int),
		params.Args["offset"].(int),
	)

	return wherePart, orderAndLimits

}
