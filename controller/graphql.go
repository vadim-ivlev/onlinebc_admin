package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	// "go/ast"
	"net/http"
	"strings"

	"onlinebc_admin/model/img"
	srv "onlinebc_admin/model/imgserver"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	gq "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// jsonStringToMap преобразует строку JSON в map[string]interface{}
func jsonStringToMap(s string) map[string]interface{} {
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(s), &m)
	return m
}

// getParamsFromBody извлекает параметры запроса из тела запроса
func getParamsFromBody(c *gin.Context) (map[string]interface{}, error) {
	r := c.Request
	mb := make(map[string]interface{})
	if r.ContentLength > 0 {
		errBodyDecode := json.NewDecoder(r.Body).Decode(&mb)
		return mb, errBodyDecode
	}
	return mb, errors.New("No body")
}

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

// getPayload3 извлекает "query", "variables", "operationName".
// Decoded body has precedence over POST over GET.
func getPayload3(c *gin.Context) (query string, variables map[string]interface{}) {

	// Проверяем на существование данных из Form Data
	query = c.PostForm("query")
	variables = jsonStringToMap(c.PostForm("variables"))

	// r := c.Request
	// // если есть тело запроса то берем из Request Payload (для Altair)
	// mb := make(map[string]interface{})
	// if r.ContentLength > 0 {
	// 	errBodyDecode := json.NewDecoder(r.Body).Decode(&mb)
	// 	if errBodyDecode == nil {
	// 		query, _ = mb["query"].(string)
	// 		variables, _ = mb["variables"].(map[string]interface{})
	// 	} else {
	// 		log.Println(errBodyDecode)
	// 	}
	// }

	// если есть тело запроса то берем из Request Payload (для Altair)
	params, errBody := getParamsFromBody(c)
	if errBody == nil {
		query, _ = params["query"].(string)
		variables, _ = params["variables"].(map[string]interface{})
	}

	return
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
		name := field.Name.Value
		if name != "__typename" {
			collect = append(collect, field.Name.Value)
		}
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

	query, variables := getPayload3(c)

	result := gq.Do(gq.Params{
		Schema:         schema,
		RequestString:  query,
		Context:        context.WithValue(context.Background(), "ginContext", c),
		VariableValues: variables,
	})

	c.JSON(200, result)
}

// SaveUploadedFile - сохраняет первый присоединенный в поле fileFieldName файл во временную директорию,
// загружает его на сервер и удаляет его из временной директории.
// Возвращает путь файла на сервере, размер файла, сообщение об ошибке.
func SaveUploadedFile(params gq.ResolveParams, fileFieldName string) (finalPath string, size int64, errMsg string) {
	// сохраняем
	filePath, size, err := img.SaveFirstFormFile(params, fileFieldName)
	if err != nil {
		return "", 0, err.Error()
	}

	// копируем всю директорию uploads_temp на сервер
	if msg := srv.CopyTempFilesToServer(); msg != "" {
		errMsg += "SaveUploadedFile.CopyTempFilesToServer: " + msg + " \n"
	}

	// удаляем файл
	if err = os.Remove(filePath); err != nil {
		errMsg += "SaveUploadedFile.Remove: " + err.Error() + " \n"
	}

	// убираем мусор (пустые директории)
	if msg := srv.RemoveEmptyDirectories(); msg != "" {
		errMsg += "SaveUploadedFile.RemoveEmptyDirectories: " + msg + " \n"
	}

	finalPath = srv.TrimLocaldir(filePath)
	return finalPath, size, errMsg
}

// SaveUploadedImage - сохраняет первый присоединенный в поле fileFieldName файл во временную директорию,
// оптимизирует его размер и порождает иконки разных размеров.
// Загружает полученные файлы на сервер и удаляет их из временной директории.
// Возвращает путь файла на сервере, ширину и высоту изображения, JSON строку иконок,  сообщение об ошибке.
func SaveUploadedImage(params gq.ResolveParams, fileFieldName string) (
	serverPath string, width int, height int, thumbsJSONStr string, errMsg string) {

	// сохраняем изображение
	filePath, _, err := img.SaveFirstFormFile(params, fileFieldName)
	if err != nil {
		return "", 0, 0, "", "SaveUploadedImage(): " + err.Error()
	}

	// проверяем допустимо ли расширение
	if !img.Params.ValidImgExtensions[strings.ToLower(filepath.Ext(filePath))] {
		// удаляем файлы вместе с директорией
		dir := path.Dir(filePath)
		if err = os.RemoveAll(dir); err != nil {
			errMsg += "SaveUploadedImage().InvalidExt.RemoveAll(): " + err.Error() + " \n"
		}
		return "", 0, 0, "", "SaveUploadedImage(): Wrong file type. Should be:" + fmt.Sprintf("%v", img.Params.ValidImgExtensions)
	}

	serverPath = srv.TrimLocaldir(filePath)

	// оптимизируем изображение
	filePath, width, height = img.OptimizeImage(filePath)

	// Генерируем иконки
	thumbsJSONStr, err = img.GenerateIcons(filePath)
	if err != nil {
		errMsg = "SaveUploadedImage(): " + err.Error()
	}

	// копируем всю директорию uploads_temp на сервер
	if msg := srv.CopyTempFilesToServer(); msg != "" {
		errMsg += "SaveUploadedImage().CopyTempFilesToServer: " + msg + " \n"
	}

	// удаляем файлы вместе с директорией
	dir := path.Dir(filePath)
	if err = os.RemoveAll(dir); err != nil {
		errMsg += "SaveUploadedImage().RemoveAll(): " + err.Error() + " \n"
	}

	// убираем мусор (пустые директории)
	if msg := srv.RemoveEmptyDirectories(); msg != "" {
		errMsg += "SaveUploadedImage.RemoveEmptyDirectories: " + msg + " \n"
	}

	return serverPath, width, height, thumbsJSONStr, errMsg
}
