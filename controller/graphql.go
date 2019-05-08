package controller

import (
	"encoding/json"
	"fmt"

	// "go/ast"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	gq "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// getFormFields извлекает имена-значения полей формы из запроса
// or builds a map with keys "query", "variables", "operationName".
// Decoded body has precedence over POST over GET.
func getPayload(r *http.Request) map[string]interface{} {
	m := make(map[string]interface{})
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	for k := range r.Form {
		m[k] = r.FormValue(k)
	}
	if r.ContentLength > 0 {
		_ = json.NewDecoder(r.Body).Decode(&m)
	}
	return m
}

var schema, _ = gq.NewSchema(gq.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

// getSelectedFields - returns list of selected fields defined in GraphQL query.
/*

	Invoke it like this:

	getSelectedFields([]string{"companies"}, resolveParams)
	// this will return []string{"id", "name"}
	In case you have a "path" you want to select from, e.g.

	query {
	a {
		b {
		x,
		y,
		z
		}
	}
	}
	Then you'd call it like this:

	getSelectedFields([]string{"a", "b"}, resolveParams)
	// Returns []string{"x", "y", "z"}

	import "github.com/graphql-go/graphql/language/ast" is added by hands.
	source: https://github.com/graphql-go/graphql/issues/125
*/
func getSelectedFields(selectionPath []string, resolveParams graphql.ResolveParams) string {
	fields := resolveParams.Info.FieldASTs
	for _, propName := range selectionPath {
		found := false
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}
				found = true
				break
			}
		}
		if !found {
			return ""
		}
	}
	var collect []string
	for _, field := range fields {
		collect = append(collect, field.Name.Value)
	}
	s := strings.Join(collect, ", ")
	return s
}

// *********************************************************************
// *********************************************************************
// *********************************************************************
// *********************************************************************
// *********************************************************************

// GraphQL исполняет GraphQL запрос
func (dummy) GraphQL(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 100*1024*1024)

	m := getPayload(c.Request)

	// Альтернативный способ. Оставлено на всякий случай
	// query, _ := c.GetPostForm("query")
	// variables, _ := c.GetPostForm("variables")

	query, _ := m["query"].(string)
	variables, _ := m["variables"].(map[string]interface{})

	result := gq.Do(gq.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: variables,
	})

	c.JSON(200, result)
}
