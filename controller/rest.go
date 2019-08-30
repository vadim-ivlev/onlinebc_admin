package controller

import (
	"fmt"
	"onlinebc_admin/model/db"
	"onlinebc_admin/model/embeds"

	"github.com/gin-gonic/gin"
)

// TODO: комментированные методы оставлены на всякий случай, если понадобится расширить API
// Общие методы************************************************************************

// LandingPage тестовая страница API в браузере.
func (dummy) LandingPage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(200, "index.html", Routes)
}


// GetFullBroadcast возвращает трансляцию с постами по её id
func (dummy) GetFullBroadcast(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_full_broadcast($1) AS res;", getIntID(c))
	jsonBytes := row["res"]
	// fmt.Fprintf(c.Writer, "%s", jsonBytes)
	amendedJsonText := embeds.AmendPostsAndAnswers(jsonBytes.([]byte))
	fmt.Fprintf(c.Writer, "%s", amendedJsonText)
}


// GetBroadcasts Получить список трансляций
func (dummy) GetBroadcasts(c *gin.Context) {
	row, _ := db.QueryRowMap("SELECT get_broadcasts() AS res;")
	fmt.Fprintf(c.Writer, "%s", row["res"])
}

