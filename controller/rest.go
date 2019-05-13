package controller

import (
	"fmt"
	"onlinebc_admin/model/db"

	"github.com/gin-gonic/gin"
)

// Общие методы************************************************************************

// LandingPage тестовая страница API в браузере.
func (dummy) LandingPage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(200, "index.html", Routes)
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

// GetFullBroadcastLegacy возвращает трансляцию с постами по её id
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
