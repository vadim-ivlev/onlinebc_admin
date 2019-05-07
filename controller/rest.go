package controller

import (
	"fmt"
	"onlinebc_admin/model/db"

	"github.com/gin-gonic/gin"
)

// ************************************************************************

// GetEntity возвращает сущность по id
func (dummy) GetEntity(c *gin.Context) {
	row, _ := db.GetRowByID(c.Param("entity"), getIntID(c))
	c.JSON(200, row)
}

// UpdateEntity обновляет сущность по id
func (dummy) UpdateEntity(c *gin.Context) {
	row, _ := db.UpdateRowByID(c.Param("entity"), getIntID(c), getPayload(c.Request))
	c.JSON(200, row)
}

// CreateEntity Создать сущность
func (dummy) CreateEntity(c *gin.Context) {
	row, _ := db.CreateRow(c.Param("entity"), getPayload(c.Request))
	c.JSON(200, row)
}

// DeleteEntity обновляет сущность по id
func (dummy) DeleteEntity(c *gin.Context) {
	row, _ := db.DeleteRowByID(c.Param("entity"), getIntID(c))
	c.JSON(200, row)
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
	row, _ := db.QueryRowMap("SELECT get_media($1) AS res;", getIntID(c))
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

// GetAnswers возвращает ответы к посту по его id
func (dummy) GetAnswers(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_answers($1) AS res;", getIntID(c))
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

// GetPosts возвращает посты трансляции по её id
func (dummy) GetPosts(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_posts($1) AS res;", getIntID(c))
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

// GetFullBroadcast возвращает трансляцию с постами по её id
func (dummy) GetFullBroadcast(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_full_broadcast($1) AS res;", getIntID(c))
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

// GetFullBroadcast возвращает трансляцию с постами по её id
func (dummy) GetFullBroadcastLegacy(c *gin.Context) {
	_ = c.Request.ParseForm()
	id := c.Request.FormValue("id")
	row, _ := db.QueryRowMap("SELECT get_full_broadcast($1) AS res;", id)
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

// GetBroadcasts Получить список трансляций
func (dummy) GetBroadcasts(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_broadcasts() AS res;")
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

// GetBroadcastList Получить список трансляций. Legacy.
func (dummy) GetBroadcastList(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_broadcasts() AS res;")
	fmt.Fprintf(c.Writer, "%s", row["res"])
}
