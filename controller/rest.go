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
	row, _ := db.QueryRowMap("SELECT get_images($1) AS res;", getIntID(c))
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

// // GetBroadcastList Получить список трансляций. Legacy.
// func (dummy) GetBroadcastList(c *gin.Context) {
// 	queryEnd := queryEndLegacy(c)
// 	row, _ := db.QueryRowMap("SELECT get_broadcast_list($1) AS res;", queryEnd)
// 	fmt.Fprintf(c.Writer, "%s", row["res"])
// }

// // queryEndLegacy возвращает вторую часть запроса на поиск трансляций.
// // Параметры запроса: ?&main=0&active=0&num=3
// // преобразует в строку: WHERE show_main_page = 0 AND is_ended = 1 LIMIT 3
// func queryEndLegacy(c *gin.Context) (res string) {
// 	show_main_page := c.Query("main")
// 	is_not_ended := c.Query("active")
// 	limit := c.Query("num")

// 	var searchConditions []string

// 	if show_main_page != "" {
// 		n, _ := strconv.Atoi(show_main_page)
// 		searchConditions = append(searchConditions, fmt.Sprintf("show_main_page = %d", n))
// 	}

// 	if is_not_ended != "" {
// 		n, _ := strconv.Atoi(is_not_ended)
// 		if n == 0 {
// 			n = 1
// 		} else {
// 			n = 0
// 		}
// 		searchConditions = append(searchConditions, fmt.Sprintf("is_ended = %d", n))
// 	}

// 	if len(searchConditions) > 0 {
// 		res = " WHERE " + strings.Join(searchConditions, " AND ")
// 	}

// 	if limit != "" {
// 		n, _ := strconv.Atoi(limit)
// 		res += fmt.Sprintf(" LIMIT %d ", n)
// 	}
// 	return res
// }
