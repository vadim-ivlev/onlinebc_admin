package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"onlinebc_admin/model/db"
	"onlinebc_admin/router"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	fmt.Println("Функциональные тесты (End to End) ******************************************************")
	readConfigs(false)
	db.WaitForDbOrExit(10)
	createDatabaseIfNotExists()
	r = router.Setup(false, false)

	flag.Parse()
	exitCode := m.Run()

	// Pretend to close our DB connection
	os.Exit(exitCode)
}

// Test_REST_GetFullBroadcast тестирование чтения Broadcast со всеми подчиненными по REST API.
func Test_REST_GetFullBroadcast(t *testing.T) {
	fmt.Println("Testing REST /api/full-broadcast/354")
	w := getNewRecorder("GET", "/api/full-broadcast/354", nil)
	assert.Equal(t, 200, w.Code)
	body := w.Body.String()
	m := jsonStringToArrayOfMaps(body)
	id := int(m[0]["id"].(float64))
	assert.Equal(t, 354, id)
}

// Test_REST_GetMedium тестирование чтения записи Medium по REST API.
func Test_REST_GetMedium(t *testing.T) {
	fmt.Println("Testing REST /get/medium/5330")
	w := getNewRecorder("GET", "/get/medium/5330", nil)
	assert.Equal(t, 200, w.Code)
	body := w.Body.String()
	m := jsonStringToMap(body)
	id := m["id"].(float64)
	assert.Equal(t, 5330., id)
}

// Test_GraphQL_GetEntityByID тестируем считывание существующих записей.
func Test_GraphQL_GetEntityByID(t *testing.T) {
	fmt.Println("Testing GraphQL query broadcast, post, medium")
	s := `
	query { 
		broadcast (id: 354) { id  title  time_created link_article }
		post(id:23952){id id_parent text author}
		medium(id:5330){id uri thumb source}
	  }	
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)

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

// Test_GraphQL_CRUD_Broadcast тестируем создание, чтение, обновление удаление записей Broadcast.
func Test_GraphQL_CRUD_Broadcast(t *testing.T) {
	// CREATE newID
	fmt.Println("Testing GraphQL mutation createBroadcast")
	s := `
	mutation {
		createBroadcast(
		  title:"new broadcast", 
		  time_created: 123, 
		  link_article:"link"
		) 
		{
		  id 
		  title 
		  time_created 
		  link_article
		}
	  }	
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	createBroadcast := data["createBroadcast"].(map[string]interface{})
	newID := int(createBroadcast["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation updateBroadcast")
	s = `
	mutation {
		updateBroadcast(
		  id: %d,
		  title:"updated broadcast", 
		  time_created: 124, 
		  link_article:"updated link2"
		) 
		{
		  id 
		  title 
		  time_created 
		  link_article
		}
	  } 
	`
	ss := fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	updateBroadcast := data["updateBroadcast"].(map[string]interface{})
	updatedTitle, ok := updateBroadcast["title"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля title")
	assert.Equal(t, "updated broadcast", updatedTitle)

	// READ rec by newID
	fmt.Println("Testing GraphQL query broadcast")
	s = `
	query { 
		broadcast (id: %d) { id  title  time_created link_article }
	  }	
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	broadcast := data["broadcast"].(map[string]interface{})
	readTitle, ok := broadcast["title"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля title")
	assert.Equal(t, "updated broadcast", readTitle)

	// DELETE rec by newID
	fmt.Println("Testing GraphQL mutation deleteBroadcast")
	s = `
	mutation {
		deleteBroadcast(
		  id: %d
		) 
		{
		  id 
		  title 
		  time_created 
		  link_article
		}
	  } 
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	deleteBroadcast := data["deleteBroadcast"].(map[string]interface{})
	deletedID := int(deleteBroadcast["id"].(float64))
	assert.Equal(t, newID, deletedID)

	// fmt.Printf("CRUD Broadcast: NewID=%d  updatedTitle ='%s' readTitle='%s' deletedID=%d \n", newID, updatedTitle, readTitle, deletedID)

}

// Test_GraphQL_CRUD_Post тестируем создание, чтение, обновление удаление записей Post.
func Test_GraphQL_CRUD_Post(t *testing.T) {

	// CREATE newID
	fmt.Println("Testing GraphQL mutation createPost")
	s := `
	mutation {
		createPost(
		  text:"new post", 
		  author: "Петров" 
		) 
		{
		  id 
		  text 
		  author 
		}
	  }	
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	createPost := data["createPost"].(map[string]interface{})
	newID := int(createPost["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation updatePost")
	s = `
	mutation {
		updatePost(
		  id: %d,
		  text:"updated post", 
		  author: "Петровский" 
		) 
		{
		  id 
		  text 
		  author 
		}
	  } 
	`
	ss := fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	updatePost := data["updatePost"].(map[string]interface{})
	updatedText, ok := updatePost["text"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля text")
	assert.Equal(t, "updated post", updatedText)

	// READ rec by newID
	fmt.Println("Testing GraphQL query post")
	s = `
	query { 
		post (id: %d) { id  text  author }
	  }	
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	post := data["post"].(map[string]interface{})
	readText, ok := post["text"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля text")
	assert.Equal(t, "updated post", readText)

	// DELETE rec by newID
	fmt.Println("Testing GraphQL mutation deletePost")
	s = `
	mutation {
		deletePost(
		  id: %d
		) 
		{
		  id 
		  text 
		  author 
		}
	  } 
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	deletePost := data["deletePost"].(map[string]interface{})
	deletedID := int(deletePost["id"].(float64))
	assert.Equal(t, newID, deletedID)

	// fmt.Printf("CRUD Post: NewID=%d  updatedText ='%s' readText='%s' deletedID=%d \n", newID, updatedText, readText, deletedID)

}

// Test_GraphQL_Upload_Images тестируем загрузку изображений.
func Test_GraphQL_Upload_Images(t *testing.T) {

	// CREATE newID
	fmt.Println("Testing GraphQL upload images")
	s := `
	mutation {

		new0: createMedium( 
			post_id: 24098, 
			source: "RT", 
			filename: "_small.gif",
			base64: "R0lGODdhBgAHAIABAAAAAP///ywAAAAABgAHAAACCoxvALfRn2JqyBQAOw=="
		) 
		{   
			id 
			post_id  
			source 
			thumb  
			uri  
		}
		
		new1: createMedium( 
			post_id: 24098, 
			source: "RT", 
			filename: "_small.png",
			base64: "iVBORw0KGgoAAAANSUhEUgAAAAYAAAAHCAIAAACk8qu6AAAALklEQVQI122NQQoAMAzCmv7/z9nBMhidFyWIotarjgHAsLTUG7qWPoj0MzR5Px5x5hf78pZ5DQAAAABJRU5ErkJggg=="
		) 
		{   
			id 
			post_id  
			source 
			thumb  
			uri  
		}

	}
	`
	// response should be
	// {
	// 	"data": {
	// 	  "new0": {
	// 		"id": 6029,
	// 		"post_id": 24098,
	// 		"source": "RT",
	// 		"thumb": "/uploads/2019/03/05/_small_thumb.gif",
	// 		"uri": "/uploads/2019/03/05/_small.gif"
	// 	  },
	// 	  "new1": {
	// 		"id": 6030,
	// 		"post_id": 24098,
	// 		"source": "RT",
	// 		"thumb": "/uploads/2019/03/05/_small_thumb.png",
	// 		"uri": "/uploads/2019/03/05/_small.png"
	// 	  }
	// 	}
	//   }

	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	new0 := data["new0"].(map[string]interface{})
	newID := int(new0["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")
	newURI := new0["uri"].(string)
	assert.Equal(t, filepath.Base(newURI), "_small.gif")
	newThumb := new0["thumb"].(string)
	assert.Equal(t, filepath.Base(newThumb), "_small_thumb.gif")

}

// Test_GraphQL_CRUD_Medium тестируем создание, чтение, обновление, удаление записей Medium.
func Test_GraphQL_CRUD_Medium(t *testing.T) {

	// CREATE newID
	fmt.Println("Testing GraphQL mutation createMedium")
	s := `
	mutation {
		createMedium(
		  post_id: 24098,
		  thumb:"new medium", 
		  source: "Петров" 
		) 
		{
		  id 
		  thumb 
		  source 
		}
	  }	
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	createMedium := data["createMedium"].(map[string]interface{})
	newID := int(createMedium["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation updateMedium")
	s = `
	mutation {
		updateMedium(
		  id: %d,
		  thumb:"updated medium", 
		  source: "Петровский" 
		) 
		{
		  id 
		  thumb 
		  source 
		}
	  } 
	`
	ss := fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	updateMedium := data["updateMedium"].(map[string]interface{})
	updatedThumb, ok := updateMedium["thumb"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля thumb")
	assert.Equal(t, "updated medium", updatedThumb)

	// READ rec by newID
	fmt.Println("Testing GraphQL query medium")
	s = `
	query { 
		medium (id: %d) { id  thumb  source }
	  }	
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	medium := data["medium"].(map[string]interface{})
	readThumb, ok := medium["thumb"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля thumb")
	assert.Equal(t, "updated medium", readThumb)

	// DELETE rec by newID
	fmt.Println("Testing GraphQL mutation deleteMedium")
	s = `
	mutation {
		deleteMedium(
		  id: %d
		) 
		{
		  id 
		  thumb 
		  source 
		}
	  } 
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	deleteMedium := data["deleteMedium"].(map[string]interface{})
	deletedID := getID(t, deleteMedium, true)
	assert.Equal(t, newID, deletedID)

	// fmt.Printf("CRUD Medium: NewID=%d  updatedThumb ='%s' readThumb='%s' deletedID=%d \n", newID, updatedThumb, readThumb, deletedID)

}

func Test_GraphQL_UpdateNONExistantID(t *testing.T) {

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation updateMedium NONEXISTANT")
	s := `
	mutation {
		updateMedium(
		  id: 777,
		  thumb:"updated medium", 
		  source: "Петровский" 
		) 
		{
		  id 
		  thumb 
		  source 
		}
	  } 
	`
	// ss := fmt.Sprintf(s, newID)
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	body := w.Body.String()
	fmt.Println(body)
	m := jsonStringToMap(body)
	data := m["data"].(map[string]interface{})
	updateMedium := data["updateMedium"].(map[string]interface{})
	updatedThumb, ok := updateMedium["thumb"].(string)
	assert.Equal(t, false, ok, "Сервер вернул нулевое значение поля Thumb")
	assert.Equal(t, "", updatedThumb)

}

func Test_GraphQL_DeleteNONExistantID(t *testing.T) {
	fmt.Println("Testing GraphQL mutation deleteMedium")
	s := `
	mutation {
		deleteMedium(
		  id: 777
		) 
		{
		  id 
		  thumb 
		  source 
		}
	  } 
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	deleteMedium := data["deleteMedium"].(map[string]interface{})
	getID(t, deleteMedium, false)
}

// ******************************************************************

func getID(t *testing.T, something map[string]interface{}, wantOk bool) int {
	floatID, ok := something["id"].(float64)
	assert.Equal(t, wantOk, ok, "Сервер вернул нулевое значение поля id")
	return int(floatID)
}

// jsonStringToMap преобразует строку JSON в map[string]interface{}
func jsonStringToMap(s string) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	return m
}

// jsonStringToArrayOfMaps преобразует строку JSON в массив []map[string]interface{}
func jsonStringToArrayOfMaps(s string) []map[string]interface{} {
	var m []map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	return m
}

func getNewRecorder(method, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Content-Length", strconv.Itoa(len(encLoad)))
	r.ServeHTTP(w, req)
	return w
}
