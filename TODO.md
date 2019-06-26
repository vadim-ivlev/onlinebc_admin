
TODO:
====

- тестирование: create_post with id_parent does not invalidate redis


- Прикрутить профилирование pprof https://habr.com/ru/company/badoo/blog/301990/


- привести в соответствие videos с onlinebc_admin 
    (docker-compose-frontend.yml, graphgl controller, ...)



```go

// getParamsFromRequest извлекает параметры запроса из *http.Request
func getParamsFromRequest(c *gin.Context) map[string]interface{} {
	r := c.Request
	m := make(map[string]interface{})
	err := r.ParseForm()
	if err != nil {
		return m
	}
	for k := range r.Form {
		m[k] = r.FormValue(k)
	}
	return m
}



```