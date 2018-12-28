package controller

import (
	"encoding/json"
	"net/http"
	"onlinebc_admin/model/db"

	"github.com/graphql-go/graphql"
)

// TYPES ****************************************************
var broadcastType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Broadcast",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.String},
		"title":          &graphql.Field{Type: graphql.String},
		"time_created":   &graphql.Field{Type: graphql.String},
		"time_begin":     &graphql.Field{Type: graphql.String},
		"is_ended":       &graphql.Field{Type: graphql.String},
		"show_date":      &graphql.Field{Type: graphql.String},
		"show_time":      &graphql.Field{Type: graphql.String},
		"is_yandex":      &graphql.Field{Type: graphql.String},
		"yandex_ids":     &graphql.Field{Type: graphql.String},
		"show_main_page": &graphql.Field{Type: graphql.String},
		"link_article":   &graphql.Field{Type: graphql.String},
		"link_img":       &graphql.Field{Type: graphql.String},
		"groups_create":  &graphql.Field{Type: graphql.String},
		"is_diary":       &graphql.Field{Type: graphql.String},
		"diary_author":   &graphql.Field{Type: graphql.String},
	},
})

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.String},
		"id_parent":    &graphql.Field{Type: graphql.String},
		"id_broadcast": &graphql.Field{Type: graphql.String},
		"text":         &graphql.Field{Type: graphql.String},
		"post_time":    &graphql.Field{Type: graphql.String},
		"post_type":    &graphql.Field{Type: graphql.String},
		"link":         &graphql.Field{Type: graphql.String},
		"has_big_img":  &graphql.Field{Type: graphql.String},
		"author":       &graphql.Field{Type: graphql.String},
	},
})

var mediumType = graphql.NewObject(graphql.ObjectConfig{
	Name: "medium",
	Fields: graphql.Fields{
		"id":      &graphql.Field{Type: graphql.String},
		"post_id": &graphql.Field{Type: graphql.String},
		"uri":     &graphql.Field{Type: graphql.String},
		"thumb":   &graphql.Field{Type: graphql.String},
		"source":  &graphql.Field{Type: graphql.String},
	},
})

// ************************************************************************

var rootQuery1 = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{

		"posts": &graphql.Field{
			Type: graphql.NewList(postType),
			Args: graphql.FieldConfigArgument{
				"id_broadcast": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// idBroadcast := params.Args["id_broadcast"].(string)
				var posts interface{}
				return posts, nil
			},
		},

		"broadcast": &graphql.Field{
			Type: broadcastType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// id := params.Args["id"].(string)
				var broadcast interface{}
				return broadcast, nil
			},
		},
	},
})

var rootMutation1 = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createBroadcast": &graphql.Field{
			Type: broadcastType,
			Args: graphql.FieldConfigArgument{
				// "id":             &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"title":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"time_created":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"time_begin":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"is_ended":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"show_date":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"show_time":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"is_yandex":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"yandex_ids":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"show_main_page": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"link_article":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"link_img":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"groups_create":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"is_diary":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"diary_author":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
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

var schema1, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery1,
	Mutation: rootMutation1,
})

// *******************************************************************************8

// GraphQL0 исполняет GraphQL запрос
func (dummy) GraphQL1(w http.ResponseWriter, r *http.Request) {
	m := getPayload(r)
	result := graphql.Do(graphql.Params{
		Schema:        schema1,
		RequestString: m["query"].(string),
	})
	json.NewEncoder(w).Encode(result)
}
