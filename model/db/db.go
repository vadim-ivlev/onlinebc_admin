package db

import (
	"database/sql"
	"fmt"
	"strings"

	//blank import

	_ "github.com/lib/pq"
)

// QueryRowResult возвращает результат запроса заданного sqlText, с возможными параметрами args.
// Применяется для исполнения запросов , INSERT, SELECT.
// Возвращает единственное значение определенное в тексте запроса.
func QueryRowResult(sqlText string, args ...interface{}) interface{} {
	conn, err := sql.Open("postgres", connectStr)
	panicIf(err)
	defer conn.Close()
	var result interface{}
	err = conn.QueryRow(sqlText, args...).Scan(&result)
	printIf(err)
	return result
}

// GetExecResult исполняет запрос заданный строкой sqlText, с возможными параметрами args.
// Применяется для исполнения запросов UPDATE, DELETE.
// sql.Result.RowsAffected() возвращает количество записей затронутых запросом.
func GetExecResult(sqlText string, args ...interface{}) sql.Result {
	conn, err := sql.Open("postgres", connectStr)
	panicIf(err)
	defer conn.Close()
	result, err1 := conn.Exec(sqlText, args...)
	printIf(err1)
	return result
}

// CreateRow Вставляет запись в таблицу tableName.
// Хэш vars задает имена и значения полей таблицы.
// Возвращает идентификатор id вставленной записи.
func CreateRow(tableName string, vars map[string]string) interface{} {
	keys, values, dollars := getKeysAndValues(vars)
	sqlText := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) RETURNING id;",
		tableName,
		strings.Join(keys, ", "),
		strings.Join(dollars, ", "))
	res := QueryRowResult(sqlText, values...)
	return res
}

// UpdateRowByID обновляет запись в таблице tableName по ее id
// полученному как значение ключа 'id' в хэше vars.
// Хэш vars задает имена и значения полей таблицы.
// Возвращает количество записей затронутых запросом.
func UpdateRowByID(tableName string, vars map[string]string) int64 {
	keys, values, dollars := getKeysAndValues(vars)
	sqlText := fmt.Sprintf("UPDATE %s SET ( %s ) = ( %s ) WHERE id = %v ;",
		tableName,
		strings.Join(keys, ", "),
		strings.Join(dollars, ", "),
		vars["id"])
	res := GetExecResult(sqlText, values...)
	num, err := res.RowsAffected()
	printIf(err)
	return num
}

// DeleteRowByID удаляет запись в таблице tableName по ее id
func DeleteRowByID(tableName string, vars map[string]string) int64 {
	sqlText := fmt.Sprintf("DELETE FROM %s WHERE id = %v ;",
		tableName,
		vars["id"])
	res := GetExecResult(sqlText)
	num, err := res.RowsAffected()
	printIf(err)
	return num
}

// getKeysAndValues возвращает срезы ключей и значений
func getKeysAndValues(vars map[string]string) ([]string, []interface{}, []string) {
	keys := []string{}
	values := make([]interface{}, 0)
	qustionMarks := []string{}
	n := 1
	for key, val := range vars {

		if val == "" {
			values = append(values, nil)
		} else {
			values = append(values, val)
		}
		keys = append(keys, key)

		qustionMarks = append(qustionMarks, fmt.Sprintf("$%v", n))
		n++
	}
	return keys, values, qustionMarks
}
