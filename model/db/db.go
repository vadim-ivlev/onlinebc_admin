package db

import (
	"database/sql"
	//blank import
	_ "github.com/lib/pq"
)

// GetJSON возвращает JSON результатов запроса заданного sqlText, с возможными параметрами.
func GetJSON(sqlText string, args ...interface{}) string {
	conn, err := sql.Open("postgres", connectStr)
	panicIf(err)
	defer conn.Close()
	var json string
	err = conn.QueryRow(sqlText, args...).Scan(&json)
	printIf(err)
	return json
}

// ExequteSQL исполняет запрос заданный строкой sqlText.
func ExequteSQL(sqlText string) error {
	conn, err := sql.Open("postgres", connectStr)
	panicIf(err)
	defer conn.Close()
	_, err1 := conn.Exec(sqlText)
	printIf(err1)
	return err1
}
