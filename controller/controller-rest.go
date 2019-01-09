package controller

import (
	"fmt"
	"onlinebc_admin/model/db"

	"github.com/gin-gonic/gin"
)

// ************************************************************************

// GetMedium возвращает медиа по id
func (dummy) GetMedium(c *gin.Context) {
	// json.NewEncoder(c.Writer).Encode(db.GetRowByID("medium", getGinIntID(c)))
	c.JSON(200, db.GetRowByID("medium", getGinIntID(c)))
}

// GetPost возвращает пост по id
func (dummy) GetPost(c *gin.Context) {
	c.JSON(200, db.GetRowByID("post", getGinIntID(c)))
}

// GetBroadcast возвращает трансляцию по id
func (dummy) GetBroadcast(c *gin.Context) {
	c.JSON(200, db.GetRowByID("broadcast", getGinIntID(c)))
}

// ************************************************************************

// UpdateMedium обновляет медиа по id
func (dummy) UpdateMedium(c *gin.Context) {
	c.JSON(200, db.UpdateRowByID("medium", getGinIntID(c), getPayload(c.Request)))
}

// UpdatePost обновляет пост по id
func (dummy) UpdatePost(c *gin.Context) {
	c.JSON(200, db.UpdateRowByID("post", getGinIntID(c), getPayload(c.Request)))
}

// UpdateBroadcast обновляет трансляцию по id
func (dummy) UpdateBroadcast(c *gin.Context) {
	c.JSON(200, db.UpdateRowByID("broadcast", getGinIntID(c), getPayload(c.Request)))
}

// ************************************************************************

// CreateMedium Создать медиа поста с идентификатором id
func (dummy) CreateMedium(c *gin.Context) {
	c.JSON(200, db.CreateRow("medium", getPayload(c.Request)))
}

// CreatePost Создать пост трансляции с идентификатором id
func (dummy) CreatePost(c *gin.Context) {
	c.JSON(200, db.CreateRow("post", getPayload(c.Request)))
}

// CreateBroadcast Создать  трансляцию
func (dummy) CreateBroadcast(c *gin.Context) {
	c.JSON(200, db.CreateRow("broadcast", getPayload(c.Request)))
}

// ************************************************************************

// DeleteMedium обновляет медиа по id
func (dummy) DeleteMedium(c *gin.Context) {
	c.JSON(200, db.DeleteRowByID("medium", getGinIntID(c)))
}

// DeletePost обновляет пост по id
func (dummy) DeletePost(c *gin.Context) {
	c.JSON(200, db.DeleteRowByID("post", getGinIntID(c)))
}

// DeleteBroadcast обновляет трансляцию по id
func (dummy) DeleteBroadcast(c *gin.Context) {
	c.JSON(200, db.DeleteRowByID("broadcast", getGinIntID(c)))
}

// ************************************************************************

// LandingPage : To test API in browser.
func (dummy) LandingPage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(200, "index.html", Routes)
}

// GetRoutes : Перечисляет доступные маршруты.  Документация API.
func (dummy) GetRoutes(c *gin.Context) {
	c.JSON(200, Routes)
}

// ************************************************************************

// GetMedia возвращает все медиа поста по его id
func (dummy) GetMedia(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_media($1) AS res;", getGinIntID(c))["res"])
}

// GetAnswers возвращает ответы к посту по его id
func (dummy) GetAnswers(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_answers($1) AS res;", getGinIntID(c))["res"])
}

// GetPosts возвращает посты трансляции по её id
func (dummy) GetPosts(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_posts($1) AS res;", getGinIntID(c))["res"])
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func (dummy) GetFullBroadcast(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_full_broadcast($1) AS res;", getGinIntID(c))["res"])
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func (dummy) GetFullBroadcastLegacy(c *gin.Context) {
	c.Request.ParseForm()
	id := c.Request.FormValue("id")
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_full_broadcast($1) AS res;", id)["res"])
}

// GetBroadcasts Получить список трансляций
func (dummy) GetBroadcasts(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_broadcasts() AS res;")["res"])
}

// GetBroadcastList Получить список трансляций. Legacy.
func (dummy) GetBroadcastList(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_broadcasts() AS res;")["res"])
}
