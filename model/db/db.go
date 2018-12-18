package db

import (
	"database/sql"
	"fmt"
	"strings"

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
func ExequteSQL(sqlText string, args ...interface{}) error {
	conn, err := sql.Open("postgres", connectStr)
	panicIf(err)
	defer conn.Close()
	_, err1 := conn.Exec(sqlText, args...)
	printIf(err1)
	return err1
}

// CreateRow Вставляет запись в таблицу tableName.
// Хэш vars задает имена и значения полей таблицы.
func CreateRow(tableName string, vars map[string]string) {
	keys, values, dollars := getKeysAndValues(vars)
	sqlText := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ;",
		tableName,
		strings.Join(keys, ", "),
		strings.Join(dollars, ", "))
	ExequteSQL(sqlText, values...)
}

// UpdateRowByID обновляет запись в таблице tableName по ее id
// полученному как значение ключа 'id' в хэше vars.
// Хэш vars задает имена и значения полей таблицы.
func UpdateRowByID(tableName string, vars map[string]string) {
	keys, values, dollars := getKeysAndValues(vars)
	sqlText := fmt.Sprintf("UPDATE %s SET ( %s ) = ( %s ) WHERE id = %v ;",
		tableName,
		strings.Join(keys, ", "),
		strings.Join(dollars, ", "),
		vars["id"])
	ExequteSQL(sqlText, values...)
}

// DeleteRowByID удаляет запись в таблице tableName по ее id
func DeleteRowByID(tableName string, vars map[string]string) {
	sqlText := fmt.Sprintf("DELETE FROM %s WHERE id = %v ;",
		tableName,
		vars["id"])
	ExequteSQL(sqlText)
}

// getKeysAndValues возвращает срезы ключей и значений
func getKeysAndValues(vars map[string]string) ([]string, []interface{}, []string) {
	keys := []string{}
	values := make([]interface{}, 0)
	qustionMarks := []string{}
	n := 1
	for key, val := range vars {
		keys = append(keys, key)
		values = append(values, val)
		qustionMarks = append(qustionMarks, fmt.Sprintf("$%v", n))
		n++
	}
	return keys, values, qustionMarks
}
