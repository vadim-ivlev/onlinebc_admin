package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gq "github.com/graphql-go/graphql"
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
