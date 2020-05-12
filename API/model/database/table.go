package database

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/louislaugier/projet-data-4A-go/database"
	"github.com/louislaugier/projet-data-4A-go/model"
)

type query struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

func (q *query) standardizeSQLQuery(values url.Values) string {
	if values.Get("limit") != "" && values.Get("offset") == "" {
		q.Limit = values.Get("limit")
		return "LIMIT " + values.Get("limit")

	}
	if values.Get("offset") != "" && values.Get("limit") == "" {
		q.Offset = values.Get("offset")
		return "OFFSET " + values.Get("offset")

	}
	if values.Get("limit") != "" && values.Get("offset") != "" {
		q.Limit = values.Get("limit")
		q.Offset = values.Get("offset")
		return "LIMIT " + values.Get("limit") + " OFFSET " + values.Get("offset")
	}
	return ""
}

// TablesGET gets all table names from database
func TablesGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rows, err := database.Db.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	for rows.Next() {
		var table string
		rows.Scan(&table)
		res.Data = append(res.Data, table)
	}
	defer rows.Close()
	res.Meta.Query = "Database table names"
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// ColsGET gets all column names for a table
func ColsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	table, _ := mux.Vars(r)["table"]
	rows, err := database.Db.Query("SELECT COLUMN_NAME FROM	information_schema.COLUMNS WHERE TABLE_NAME = '" + table + "';")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	for rows.Next() {
		var column string
		rows.Scan(&column)
		res.Data = append(res.Data, column)
	}
	defer rows.Close()
	res.Meta.Query = "Table column names for " + table
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}

// RowsGET gets each record from a table
func RowsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	table, _ := mux.Vars(r)["table"]
	var q query
	u := r.URL.Query()
	s := q.standardizeSQLQuery(u)
	rows, err := database.Db.Query("SELECT * FROM " + table + " " + s + ";")
	res := model.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "OK",
	}
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "Database Error"
		log.Println("Error with transaction:", err)
	}
	cols, _ := rows.Columns()
	for rows.Next() {
		pointers := make([]interface{}, len(cols))
		container := make([]string, len(cols))
		for i := range pointers {
			pointers[i] = &container[i]
		}
		rows.Scan(pointers...)
		res.Data = append(res.Data, container)
	}
	defer rows.Close()
	res.Meta.Query = q
	res.Meta.ResultCount = len(res.Data)
	json.NewEncoder(w).Encode(res)
}
