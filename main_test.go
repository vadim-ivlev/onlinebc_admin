package main

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
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
	db.ExitIfNoDB()
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

func TestGetMedium1(t *testing.T) {
	w := getRecorder("GET", "/get/medium/5330", nil)

	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	m := jsonStringToMap(body)
	id := m["id"].(float64)

	assert.Equal(t, 5330., id)

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
	r.ServeHTTP(w, req)
	return w
}
