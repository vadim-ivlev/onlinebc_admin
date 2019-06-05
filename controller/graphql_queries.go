package controller

import (
	"fmt"
	"onlinebc_admin/model/db"
	"strings"

	gq "github.com/graphql-go/graphql"
)

// ************************************************************************

var rootQuery = gq.NewObject(gq.ObjectConfig{
	Name: "Query",
	Fields: gq.Fields{

		"get_broadcast": &gq.Field{
			Type:        fullBroadcastType,
			Description: "Показать трансляцию по идентификатору c постами, ответами и медиа",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор трансляции",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				fields := getSelectedFields([]string{"get_broadcast"}, params)
				return db.QueryRowMap("SELECT "+fields+" FROM full_broadcast WHERE id = $1 ;", params.Args["id"].(int))
			},
		},

		"get_post": &gq.Field{
			Type:        fullPostType,
			Description: "Показать пост с ответами и изображениями по идентификатору поста",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор поста",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				fields := getSelectedFields([]string{"get_post"}, params)
				return db.QueryRowMap("SELECT "+fields+" FROM full_post WHERE id = $1 ;", params.Args["id"].(int))
			},
		},

		"get_image": &gq.Field{
			Type:        imageType,
			Description: "Показать изображение по идентификатору",
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор изображения",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				fields := getSelectedFields([]string{"get_image"}, params)
				return db.QueryRowMap("SELECT "+fields+" FROM image WHERE id = $1 ;", params.Args["id"].(int))
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
				fields := getSelectedFields([]string{"get_broadcast_posts"}, params)
				return db.QuerySliceMap("SELECT "+fields+" FROM post WHERE id_broadcast = $1 "+showAnswersCondition+" ORDER BY post_time DESC ;", params.Args["id_broadcast"].(int))
			},
		},

		"get_post_images": &gq.Field{
			Type:        gq.NewList(imageType),
			Description: "Получить фотографии поста по его идентификатору.",
			Args: gq.FieldConfigArgument{
				"post_id": &gq.ArgumentConfig{
					Type:        gq.NewNonNull(gq.Int),
					Description: "Идентификатор поста",
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				fields := getSelectedFields([]string{"get_post_images"}, params)
				return db.QuerySliceMap("SELECT "+fields+" FROM image WHERE post_id = $1 ;", params.Args["post_id"].(int))
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
				fields := getSelectedFields([]string{"get_post_answers"}, params)
				return db.QuerySliceMap("SELECT "+fields+" FROM post WHERE id_parent = $1 ;", params.Args["id_parent"].(int))
			},
		},

		"list_broadcast": &gq.Field{
			Type:        listBroadcastType,
			Description: "Получить список трансляций и их количество.",
			Args: gq.FieldConfigArgument{
				"search": &gq.ArgumentConfig{
					Type:        gq.String,
					Description: "Строка полнотекстового поиска.",
					// DefaultValue: "",
				},

				"id": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор трансляции",
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
					Description: "1 если трансляция закончена, 0 - если нет.",
					// DefaultValue: 1,
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
				fields := getSelectedFields([]string{"list_broadcast", "list"}, params)

				list, err := db.QuerySliceMap("SELECT " + fields + " FROM broadcast" + wherePart + orderAndLimits)
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
					Type:        gq.String,
					Description: "Строка полнотекстового поиска.",
					// DefaultValue: "",
				},

				"id": &gq.ArgumentConfig{
					Type:        gq.Int,
					Description: "Идентификатор трансляции",
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
					Description: "1 если трансляция закончена, 0 - если нет.",
					// DefaultValue: 1,
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
				fields := getSelectedFields([]string{"list_full_broadcast", "list"}, params)

				list, err := db.QuerySliceMap("SELECT " + fields + " FROM full_broadcast" + wherePart + orderAndLimits)
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

	var searchConditions []string

	search, ok := params.Args["search"].(string)
	if ok && len(search) > 0 {
		searchConditions = append(searchConditions,
			fmt.Sprintf("to_tsvector('russian', title) @@ plainto_tsquery('russian','%s') ", search))
	}

	is_ended, ok := params.Args["is_ended"].(int)
	if ok {
		searchConditions = append(searchConditions,
			fmt.Sprintf("is_ended = %d ", is_ended))
	}

	if len(searchConditions) > 0 {
		wherePart = " WHERE " + strings.Join(searchConditions, " AND ")
	}

	orderAndLimits = fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d ;",
		params.Args["order"].(string),
		params.Args["limit"].(int),
		params.Args["offset"].(int),
	)

	return wherePart, orderAndLimits

}
