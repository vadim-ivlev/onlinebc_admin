package main

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"onlinebc_admin/model/db"
	"onlinebc_admin/router"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	readConfigs(false)
	db.WaitForDbOrExit(10)
	createDatabaseIfNotExists()
	r = router.SetupRouter(false)

	flag.Parse()
	exitCode := m.Run()

	// Pretend to close our DB connection
	os.Exit(exitCode)
}

func TestGetMedium(t *testing.T) {
	w := getRecorder("GET", "/get/medium/5330", nil)

	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	m := jsonStringToMap(body)
	id := m["id"].(float64)

	assert.Equal(t, 5330., id)

}

func Test_GraphQL_GetEntityByID(t *testing.T) {
	s := `
	query { 
		broadcast (id: 354) { id  title  time_created link_article }
		post(id:23952){id id_parent text author}
		medium(id:5330){id uri thumb source}
	  }	
	`
	w := getRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)

	// uncomment for POST request
	// load := url.Values{}
	// load.Set("query", s)
	// encLoad := load.Encode()
	// w := getRecorder("POST", "/graphql", strings.NewReader(encLoad))

	assert.Equal(t, 200, w.Code)

	body := w.Body.String()

	m := jsonStringToMap(body)

	data := m["data"].(map[string]interface{})
	broadcast := data["broadcast"].(map[string]interface{})
	post := data["post"].(map[string]interface{})
	medium := data["medium"].(map[string]interface{})

	assert.Equal(t, 354., broadcast["id"].(float64))
	assert.Equal(t, 23952., post["id"].(float64))
	assert.Equal(t, 5330., medium["id"].(float64))

}

// ******************************************************************

func jsonStringToMap(s string) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	return m
}

func jsonStringToArray(s string) []map[string]interface{} {
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		panic(err)
	}
	return a
}

func getRecorder(method, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Content-Length", strconv.Itoa(len(encLoad)))
	r.ServeHTTP(w, req)
	return w
}
