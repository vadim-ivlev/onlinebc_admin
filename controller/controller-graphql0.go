package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

// ************************************************************************
func Filter(songs []Song, f func(Song) bool) []Song {
	vsf := make([]Song, 0)
	for _, v := range songs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

var rootQuery0 = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"songs": &graphql.Field{
			Type: graphql.NewList(songType),
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	return songs, nil
			// },

			Args: graphql.FieldConfigArgument{
				// "album": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"album": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				songsFields := getSelectedFields([]string{"songs"}, params)
				fmt.Println(songsFields)

				album := params.Args["album"]
				if album == nil {
					s := `
					[
						{
							"album": "ts-fearless",
							"title": "Fearless"
						},
						{
							"album": "ts-fearless",
							"title": "Fifteen"
						}
						]
					`
					var r []map[string]interface{}
					err := json.Unmarshal([]byte(s), &r)
					if err != nil {
						println(err.Error())
					}
					return r, nil

				}
				filtered := Filter(songs, func(v Song) bool {
					return strings.Contains(v.Album, album.(string))
				})
				return filtered, nil
			},
		},

		"album": &graphql.Field{
			Type: albumType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				for _, album := range albums {
					if album.ID == id {
						return album, nil
					}
				}
				return nil, nil
			},
		},
	},
})

var rootMutation0 = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{

		"createSong": &graphql.Field{
			Type: songType,
			Args: graphql.FieldConfigArgument{
				"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"album":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"title":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"duration": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var song Song
				song.ID = params.Args["id"].(string)
				song.Album = params.Args["album"].(string)
				song.Title = params.Args["title"].(string)
				song.Duration = params.Args["duration"].(string)
				songs = append(songs, song)
				return song, nil
			},
		},
		// =====================================================

	},
})

var schema0, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery0,
	Mutation: rootMutation0,
})

// *******************************************************************************8

// GraphQL0 исполняет GraphQL запрос
func (dummy) GraphQL0(w http.ResponseWriter, r *http.Request) {
	m := getPayload(r)
	result := graphql.Do(graphql.Params{
		Schema:        schema0,
		RequestString: m["query"].(string),
	})
	json.NewEncoder(w).Encode(result)
}
