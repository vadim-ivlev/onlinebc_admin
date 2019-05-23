
TODO:
====


- Прикрутить профилирование pprof https://habr.com/ru/company/badoo/blog/301990/


- Убрать  graphql тесты из routes.yaml   (svelte)(gqtest)
- перенести функциональные тесты из main_test.go на junit.
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