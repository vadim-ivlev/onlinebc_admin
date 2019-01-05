package db

import (
	"database/sql"
	"fmt"
	"strings"

	//blank import

	"github.com/jmoiron/sqlx"
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

// QueryRowMap возвращает результат запроса заданного sqlText, с возможными параметрами args.
// Применяется для исполнения запросов , INSERT, SELECT.
// Возвращает map[string]interface{}.
func QueryRowMap(sqlText string, args ...interface{}) map[string]interface{} {
	conn, err := sqlx.Open("postgres", connectStr)
	panicIf(err)
	defer conn.Close()
	result := make(map[string]interface{})
	err = conn.QueryRowx(sqlText, args...).MapScan(result)
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
// fieldValues задает имена и значения полей таблицы.
// Возвращает map[string]interface{} новой записи таблицы.
func CreateRow(tableName string, fieldValues map[string]interface{}) map[string]interface{} {
	keys, values, dollars := getKeysAndValues(fieldValues)
	sqlText := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) RETURNING * ;",
		tableName, strings.Join(keys, ", "), strings.Join(dollars, ", "))
	res := QueryRowMap(sqlText, values...)
	return res
}

// UpdateRowByID обновляет запись в таблице tableName по ее id.
// map fieldValues задает имена и значения полей таблицы.
// Возвращает map[string]interface{} обновленной записи таблицы.
func UpdateRowByID(tableName string, id int, fieldValues map[string]interface{}) map[string]interface{} {
	keys, values, dollars := getKeysAndValues(fieldValues)
	sqlText := fmt.Sprintf("UPDATE %s SET ( %s ) = ( %s ) WHERE id = %v RETURNING * ;",
		tableName, strings.Join(keys, ", "), strings.Join(dollars, ", "), id)
	res := QueryRowMap(sqlText, values...)
	return res
}

// DeleteRowByID удаляет запись в таблице tableName по ее id.
// Возвращает map[string]interface{} удаленной записи таблицы.
func DeleteRowByID(tableName string, id int) map[string]interface{} {
	sqlText := fmt.Sprintf("DELETE FROM %s WHERE id = %v RETURNING * ;", tableName, id)
	res := QueryRowMap(sqlText)
	return res
}

// getKeysAndValues возвращает срезы ключей, значений и символов доллара $n.
func getKeysAndValues(vars map[string]interface{}) ([]string, []interface{}, []string) {
	keys := []string{}
	values := make([]interface{}, 0)
	dollars := []string{}
	n := 1
	for key, val := range vars {
		// TODO: обработка пустых значений числовых полей. Сделать типы полей форм.
		if val == "" {
			values = append(values, nil)
		} else {
			values = append(values, val)
		}
		// values = append(values, val)
		keys = append(keys, key)
		dollars = append(dollars, fmt.Sprintf("$%v", n))
		n++
	}
	return keys, values, dollars
}
