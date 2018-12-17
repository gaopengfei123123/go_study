package main

import (
	"database/sql"
	"fmt"
	_ "github.com/GO-SQL-Driver/MySQL"
	"net/http"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123123@tcp(localhost:33060)/go?charset=utf8")
	// 设置最大连接数
	db.SetMaxOpenConns(300)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(100)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(time.Second * 5)
	err := db.Ping()
	checkErr(err)

}

func main() {
	startServer(":9999")
}

func startServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		record := doSomething()
		fmt.Fprintln(w, record)
	})

	fmt.Printf("listening http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	checkErr(err)
}

func doSomething() map[string]string {
	rows, err := db.Query("SELECT * FROM test LIMIT 1")
	checkErr(err)
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	// fmt.Println(record)

	return record
}

func doSomething2() map[string]string {
	dbConn, _ := sql.Open("mysql", "root:123123@tcp(localhost:33060)/go?charset=utf8")
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT * FROM test LIMIT 1")
	checkErr(err)
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	// fmt.Println(record)

	return record
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
