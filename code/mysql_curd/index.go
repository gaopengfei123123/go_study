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
	defer db.Close()
	checkErr(err)

	//insert
	stmt, err := db.Prepare("INSERT test SET name=? , age=? , created_at=?")
	checkErr(err)
	res, err := stmt.Exec("gaopengfei", 22, time.Now().Unix())
	checkErr(err)
	//获取插入数据的 id
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("insert id %d \n", id)

	//update
	stmt, err = db.Prepare("UPDATE test SET age=? WHERE id < ?")
	checkErr(err)
	res, err = stmt.Exec(233, 4)
	checkErr(err)
	//输出受影响的条数
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("affected num is %d \n", affect)

	//select
	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)
	for rows.Next() {
		var id int
		var name string
		var age int
		var createdAt int //go 不让使用下划线的方式定义变量,如果字段是 created_at 需要写成 createdAt 驼峰命名法
		err = rows.Scan(&id, &name, &age, &createdAt)
		checkErr(err)
		//将int 转成 int64格式,方便格式化时间
		var tm64 int64
		tm64 = int64(createdAt)
		tm := time.Unix(tm64, 0)
		fmt.Printf("id is %d; name is %s; age is %d: created_at is %s \n", id, name, age, tm.Format("2006-01-02 03:04:05 PM"))
	}

	//delete
	stmt, err = db.Prepare("DELETE FROM test WHERE id=?")
	checkErr(err)
	res, err = stmt.Exec(3)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Printf("affected num is %d \n", affect)
}

// 操作的数据表结构
// create table test(
// 	id int primary key auto_increment,
// 	name varchar(20) default '',
// 	age int default 0,
// 	created_at int default 0
// )
