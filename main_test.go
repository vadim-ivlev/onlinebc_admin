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
	db.CreateDatabaseIfNotExists()
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

// Test_GraphQL_GetEntityByID тестируем считывание существующих записей.
func Test_GraphQL_GetEntityByID(t *testing.T) {
	fmt.Println("Testing GraphQL query get_broadcast, post, image")
	s := `
	query { 
		get_broadcast (id: 354) { id  title  time_created link_article }
		get_post(id:23952){id id_parent text author}
		get_image(id:5330){id filepath source}
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
	get_broadcast := data["get_broadcast"].(map[string]interface{})
	get_post := data["get_post"].(map[string]interface{})
	get_image := data["get_image"].(map[string]interface{})

	assert.Equal(t, 354., get_broadcast["id"].(float64))
	assert.Equal(t, 23952., get_post["id"].(float64))
	assert.Equal(t, 5330., get_image["id"].(float64))

}

// Test_GraphQL_CRUD_Broadcast тестируем создание, чтение, обновление удаление записей Broadcast.
func Test_GraphQL_CRUD_Broadcast(t *testing.T) {
	// CREATE newID
	fmt.Println("Testing GraphQL mutation create_broadcast")
	s := `
	mutation {
		create_broadcast(
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
	create_broadcast := data["create_broadcast"].(map[string]interface{})
	newID := int(create_broadcast["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation update_broadcast")
	s = `
	mutation {
		update_broadcast(
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
	update_broadcast := data["update_broadcast"].(map[string]interface{})
	updatedTitle, ok := update_broadcast["title"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля title")
	assert.Equal(t, "updated broadcast", updatedTitle)

	// READ rec by newID
	fmt.Println("Testing GraphQL query broadcast")
	s = `
	query { 
		get_broadcast (id: %d) { id  title  time_created link_article }
	  }	
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	get_broadcast := data["get_broadcast"].(map[string]interface{})
	readTitle, ok := get_broadcast["title"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля title")
	assert.Equal(t, "updated broadcast", readTitle)

	// DELETE rec by newID
	fmt.Println("Testing GraphQL mutation delete_broadcast")
	s = `
	mutation {
		delete_broadcast(
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
	delete_broadcast := data["delete_broadcast"].(map[string]interface{})
	deletedID := int(delete_broadcast["id"].(float64))
	assert.Equal(t, newID, deletedID)

	// fmt.Printf("CRUD Broadcast: NewID=%d  updatedTitle ='%s' readTitle='%s' deletedID=%d \n", newID, updatedTitle, readTitle, deletedID)

}

// Test_GraphQL_CRUD_Post тестируем создание, чтение, обновление удаление записей Post.
func Test_GraphQL_CRUD_Post(t *testing.T) {

	// CREATE newID
	fmt.Println("Testing GraphQL mutation create_post")
	s := `
	mutation {
		create_post(
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
	create_post := data["create_post"].(map[string]interface{})
	newID := int(create_post["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation update_post")
	s = `
	mutation {
		update_post(
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
	update_post := data["update_post"].(map[string]interface{})
	updatedText, ok := update_post["text"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля text")
	assert.Equal(t, "updated post", updatedText)

	// READ rec by newID
	fmt.Println("Testing GraphQL query post")
	s = `
	query { 
		get_post (id: %d) { id  text  author }
	  }	
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	get_post := data["get_post"].(map[string]interface{})
	readText, ok := get_post["text"].(string)
	assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля text")
	assert.Equal(t, "updated post", readText)

	// DELETE rec by newID
	fmt.Println("Testing GraphQL mutation delete_post")
	s = `
	mutation {
		delete_post(
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
	delete_post := data["delete_post"].(map[string]interface{})
	deletedID := int(delete_post["id"].(float64))
	assert.Equal(t, newID, deletedID)

	// fmt.Printf("CRUD Post: NewID=%d  updatedText ='%s' readText='%s' deletedID=%d \n", newID, updatedText, readText, deletedID)

}

// Test_GraphQL_Upload_Images тестируем загрузку изображений.
func Test_GraphQL_Upload_Images(t *testing.T) {

	// CREATE newID
	fmt.Println("Testing GraphQL upload images")
	s := `
	mutation {

		new0: create_image( 
			post_id: 24098, 
			source: "RT", 
			filename: "_small.gif",
			base64: "R0lGODdhBgAHAIABAAAAAP///ywAAAAABgAHAAACCoxvALfRn2JqyBQAOw=="
		) 
		{   
			id 
			post_id  
			source 
			thumbs {
				type
				filepath
				height
				width
			}
			filepath  
		}
		
		new1: create_image( 
			post_id: 24098, 
			source: "RT", 
			filename: "_small.png",
			base64: "iVBORw0KGgoAAAANSUhEUgAAAAYAAAAHCAIAAACk8qu6AAAALklEQVQI122NQQoAMAzCmv7/z9nBMhidFyWIotarjgHAsLTUG7qWPoj0MzR5Px5x5hf78pZ5DQAAAABJRU5ErkJggg=="
		) 
		{   
			id 
			post_id  
			source 
			thumbs {
				type
				filepath
				height
				width
			}
			filepath  
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
	// 		"filepath": "/uploads/2019/03/05/_small.gif"
	// 	  },
	// 	  "new1": {
	// 		"id": 6030,
	// 		"post_id": 24098,
	// 		"source": "RT",
	// 		"thumb": "/uploads/2019/03/05/_small_thumb.png",
	// 		"filepath": "/uploads/2019/03/05/_small.png"
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
	newURI := new0["filepath"].(string)
	assert.Equal(t, filepath.Base(newURI), "_small.gif")
	// newThumb := new0["thumb"].(string)
	// assert.Equal(t, filepath.Base(newThumb), "_small_thumb.gif")

}

// Test_GraphQL_CRUD_Image тестируем создание, чтение, обновление, удаление записей Image.
func Test_GraphQL_CRUD_Image(t *testing.T) {

	// CREATE newID
	fmt.Println("Testing GraphQL mutation create_image")
	s := `
	mutation {
		create_image(
		  post_id: 24098,
		  source: "Петров" 
		) 
		{
		  id 
		  thumbs {
			type
			filepath
			height
			width
			}
	  	  source 
		}
	  }	
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	create_image := data["create_image"].(map[string]interface{})
	newID := int(create_image["id"].(float64))
	assert.True(t, newID > 0, "New ID must be greater than 0")

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation update_image")
	s = `
	mutation {
		update_image(
		  id: %d,
		  source: "Петровский" 
		) 
		{
		  id 
		  thumbs {
			type
			filepath
			height
			width
		}
	  source 
		}
	  } 
	`
	ss := fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	// update_image := data["update_image"].(map[string]interface{})
	// updatedThumb, ok := update_image["thumb"].(string)
	// assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля thumb")
	// assert.Equal(t, "updated get_image", updatedThumb)

	// READ rec by newID
	fmt.Println("Testing GraphQL query image")
	s = `
	query { 
		get_image (id: %d) { id  thumb  source }
	  }	
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	// get_image := data["get_image"].(map[string]interface{})
	// readThumb, ok := get_image["thumb"].(string)
	// assert.Equal(t, true, ok, "Сервер вернул нулевое значение поля thumb")
	// assert.Equal(t, "updated image", readThumb)

	// DELETE rec by newID
	fmt.Println("Testing GraphQL mutation delete_image")
	s = `
	mutation {
		delete_image(
		  id: %d
		) 
		{
		  id 
		  thumbs {
				type
				filepath
				height
				width
			}
		  source 
		}
	  } 
	`
	ss = fmt.Sprintf(s, newID)
	w = getNewRecorder("GET", "/graphql?query="+url.QueryEscape(ss), nil)
	assert.Equal(t, 200, w.Code)
	m = jsonStringToMap(w.Body.String())
	data = m["data"].(map[string]interface{})
	delete_image := data["delete_image"].(map[string]interface{})
	deletedID := getID(t, delete_image, true)
	assert.Equal(t, newID, deletedID)

	// fmt.Printf("CRUD Image: NewID=%d  updatedThumb ='%s' readThumb='%s' deletedID=%d \n", newID, updatedThumb, readThumb, deletedID)

}

func Test_GraphQL_UpdateNONExistantID(t *testing.T) {

	// UPDATE rec by newID
	fmt.Println("Testing GraphQL mutation update_image NONEXISTANT")
	s := `
	mutation {
		update_image(
		  id: 777,
		  source: "Петровский" 
		) 
		{
		  id 
		  thumbs {
				type
				filepath
				height
				width
			}
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
	update_image := data["update_image"].(map[string]interface{})
	updatedThumb, ok := update_image["thumb"].(string)
	assert.Equal(t, false, ok, "Сервер вернул нулевое значение поля Thumb")
	assert.Equal(t, "", updatedThumb)

}

func Test_GraphQL_DeleteNONExistantID(t *testing.T) {
	fmt.Println("Testing GraphQL mutation delete_image")
	s := `
	mutation {
		delete_image(
		  id: 777
		) 
		{
		  id 
	      thumbs {
				type
				filepath
				height
				width
			}
		  source 
		}
	  } 
	`
	w := getNewRecorder("GET", "/graphql?query="+url.QueryEscape(s), nil)
	assert.Equal(t, 200, w.Code)
	m := jsonStringToMap(w.Body.String())
	data := m["data"].(map[string]interface{})
	delete_image := data["delete_image"].(map[string]interface{})
	getID(t, delete_image, false)
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
