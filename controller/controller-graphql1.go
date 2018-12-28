package controller

import (
	"encoding/json"
	"net/http"
	"onlinebc_admin/model/db"

	gq "github.com/graphql-go/graphql"
)

// TYPES ****************************************************
var broadcastType = gq.NewObject(gq.ObjectConfig{
	Name: "Broadcast",
	Fields: gq.Fields{
		"id":             &gq.Field{Type: gq.String},
		"title":          &gq.Field{Type: gq.String},
		"time_created":   &gq.Field{Type: gq.String},
		"time_begin":     &gq.Field{Type: gq.String},
		"is_ended":       &gq.Field{Type: gq.String},
		"show_date":      &gq.Field{Type: gq.String},
		"show_time":      &gq.Field{Type: gq.String},
		"is_yandex":      &gq.Field{Type: gq.String},
		"yandex_ids":     &gq.Field{Type: gq.String},
		"show_main_page": &gq.Field{Type: gq.String},
		"link_article":   &gq.Field{Type: gq.String},
		"link_img":       &gq.Field{Type: gq.String},
		"groups_create":  &gq.Field{Type: gq.String},
		"is_diary":       &gq.Field{Type: gq.String},
		"diary_author":   &gq.Field{Type: gq.String},
	},
})

var postType = gq.NewObject(gq.ObjectConfig{
	Name: "Post",
	Fields: gq.Fields{
		"id":           &gq.Field{Type: gq.String},
		"id_parent":    &gq.Field{Type: gq.String},
		"id_broadcast": &gq.Field{Type: gq.String},
		"text":         &gq.Field{Type: gq.String},
		"post_time":    &gq.Field{Type: gq.String},
		"post_type":    &gq.Field{Type: gq.String},
		"link":         &gq.Field{Type: gq.String},
		"has_big_img":  &gq.Field{Type: gq.String},
		"author":       &gq.Field{Type: gq.String},
	},
})

var mediumType = gq.NewObject(gq.ObjectConfig{
	Name: "medium",
	Fields: gq.Fields{
		"id":      &gq.Field{Type: gq.String},
		"post_id": &gq.Field{Type: gq.String},
		"uri":     &gq.Field{Type: gq.String},
		"thumb":   &gq.Field{Type: gq.String},
		"source":  &gq.Field{Type: gq.String},
	},
})

// ************************************************************************

var rootQuery1 = gq.NewObject(gq.ObjectConfig{
	Name: "Query",
	Fields: gq.Fields{

		"posts": &gq.Field{
			Type: gq.NewList(postType),
			Args: gq.FieldConfigArgument{
				"id_broadcast": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				// idBroadcast := params.Args["id_broadcast"].(string)

				var posts interface{}
				return posts, nil
			},
		},

		"broadcast": &gq.Field{
			Type: broadcastType,
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				// selectedFields := getSelectedFields([]string{"broadcast"}, params)
				// fieldList := strings.Join(selectedFields, ", ")
				broadcast := db.QueryRowResult1("SELECT * FROM broadcast WHERE id=$1;", id)
				// broadcast := db.QueryRowResult1("SELECT $1 FROM broadcast WHERE id=$2;", fieldList, id)
				// var broadcast interface{}
				return broadcast, nil
			},
		},
	},
})

var rootMutation1 = gq.NewObject(gq.ObjectConfig{
	Name: "Mutation",
	Fields: gq.Fields{
		"createBroadcast": &gq.Field{
			Type: broadcastType,
			Args: gq.FieldConfigArgument{
				// "id":             &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"title":          &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"time_created":   &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"time_begin":     &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"is_ended":       &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"show_date":      &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"show_time":      &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"is_yandex":      &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"yandex_ids":     &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"show_main_page": &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"link_article":   &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"link_img":       &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"groups_create":  &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"is_diary":       &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
				"diary_author":   &gq.ArgumentConfig{Type: gq.NewNonNull(gq.String)},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				var m = make(map[string]string)
				// m["id"] = params.Args["id"].(string)
				m["title"] = params.Args["title"].(string)
				m["time_created"] = params.Args["time_created"].(string)
				m["time_begin"] = params.Args["time_begin"].(string)
				m["is_ended"] = params.Args["is_ended"].(string)
				m["show_date"] = params.Args["show_date"].(string)
				m["show_time"] = params.Args["show_time"].(string)
				m["is_yandex"] = params.Args["is_yandex"].(string)
				m["yandex_ids"] = params.Args["yandex_ids"].(string)
				m["show_main_page"] = params.Args["show_main_page"].(string)
				m["link_article"] = params.Args["link_article"].(string)
				m["link_img"] = params.Args["link_img"].(string)
				m["groups_create"] = params.Args["groups_create"].(string)
				m["is_diary"] = params.Args["is_diary"].(string)
				m["diary_author"] = params.Args["diary_author"].(string)
				db.CreateRow("broadcast", m)
				return m, nil
			},
		},
		// =====================================================

	},
})

var schema1, _ = gq.NewSchema(gq.SchemaConfig{
	Query:    rootQuery1,
	Mutation: rootMutation1,
})

// *******************************************************************************8

// GraphQL0 исполняет GraphQL запрос
func (dummy) GraphQL1(w http.ResponseWriter, r *http.Request) {
	m := getPayload(r)
	result := gq.Do(gq.Params{
		Schema:        schema1,
		RequestString: m["query"].(string),
	})
	json.NewEncoder(w).Encode(result)
}
