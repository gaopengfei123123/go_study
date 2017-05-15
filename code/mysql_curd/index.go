package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/GO-SQL-Driver/MySQL"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	db, err := sql.Open("mysql", "root:123123@tcp(localhost:33060)/go?charset=utf8")
	checkErr(err)
	//insert
	stmt, err := db.Prepare("INSERT test SET name=? , age=? , created_at=?")
	checkErr(err)

	res, err := stmt.Exec("gaopengfei", 22, time.Now().Unix())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

// 操作的数据表结构
// create table test(
// 	id int primary key auto_increment,
// 	name varchar(20) default '',
// 	age int default 0,
// 	created_at int default 0
// )
