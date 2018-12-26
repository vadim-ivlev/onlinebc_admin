package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// TYPES ****************************************************

type Album struct {
	ID     string `json:"id,omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Type   string `json:"type"`
}

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

// DATA *******************************************************

var albums []Album = []Album{
	Album{
		ID:     "ts-fearless",
		Artist: "1",
		Title:  "Fearless",
		Year:   "2008",
		Type:   "album",
	},
}

var artists []Artist = []Artist{
	Artist{
		ID:   "1",
		Name: "Taylor Swift",
		Type: "artist",
	},
}

var songs []Song = []Song{
	Song{
		ID:       "1",
		Album:    "ts-fearless",
		Title:    "Fearless",
		Duration: "4:01",
		Type:     "song",
	},
	Song{
		ID:       "2",
		Album:    "ts-fearless",
		Title:    "Fifteen",
		Duration: "4:54",
		Type:     "song",
	},
}

// *******************************************

var songType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"album":    &graphql.Field{Type: graphql.String},
		"title":    &graphql.Field{Type: graphql.String},
		"duration": &graphql.Field{Type: graphql.String},
	},
})

var artistType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Artist",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"type": &graphql.Field{Type: graphql.String},
	},
})

var albumType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Album",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.String},
		"artist": &graphql.Field{Type: graphql.String},
		"title":  &graphql.Field{Type: graphql.String},
		"year":   &graphql.Field{Type: graphql.String},
		"genre":  &graphql.Field{Type: graphql.String},
		"type":   &graphql.Field{Type: graphql.String},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"songs": &graphql.Field{
			Type: graphql.NewList(songType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return songs, nil
			},
		},
	},
})

// schema, _ := graphql.NewSchema(graphql.SchemaConfig{})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

// GraphQLHandler выполняет graphql запрос
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		fmt.Println(err)
	}
	r.ParseForm()
	result := graphql.Do(graphql.Params{
		Schema: schema,
		// RequestString: r.URL.Query().Get("query"),
		RequestString: r.FormValue("query"),
	})
	json.NewEncoder(w).Encode(result)
}
