package controller

import (
	"fmt"
	"onlinebc_admin/model/db"

	"github.com/gin-gonic/gin"
)

// ************************************************************************

// GetEntity возвращает сущность по id
func (dummy) GetEntity(c *gin.Context) {
	c.JSON(200, db.GetRowByID(c.Param("entity"), getIntID(c)))
}

// UpdateEntity обновляет сущность по id
func (dummy) UpdateEntity(c *gin.Context) {
	c.JSON(200, db.UpdateRowByID(c.Param("entity"), getIntID(c), getPayload(c.Request)))
}

// CreateEntity Создать сущность
func (dummy) CreateEntity(c *gin.Context) {
	c.JSON(200, db.CreateRow(c.Param("entity"), getPayload(c.Request)))
}

// DeleteEntity обновляет сущность по id
func (dummy) DeleteEntity(c *gin.Context) {
	c.JSON(200, db.DeleteRowByID(c.Param("entity"), getIntID(c)))
}

// Общие методы************************************************************************

// LandingPage тестовая страница API в браузере.
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
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_media($1) AS res;", getIntID(c))["res"])
}

// GetAnswers возвращает ответы к посту по его id
func (dummy) GetAnswers(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_answers($1) AS res;", getIntID(c))["res"])
}

// GetPosts возвращает посты трансляции по её id
func (dummy) GetPosts(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_posts($1) AS res;", getIntID(c))["res"])
}

// GetFullBroadcast возвращает трасляцию с постами по её id
func (dummy) GetFullBroadcast(c *gin.Context) {
	fmt.Fprintf(c.Writer, "%s", db.QueryRowMap("SELECT get_full_broadcast($1) AS res;", getIntID(c))["res"])
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
