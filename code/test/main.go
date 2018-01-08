package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:123123@tcp(127.0.0.1:3306)/go?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(User))

    // create table
    // orm.RunSyncdb("default", false, true)
}

func main() {
    db := orm.NewOrm()
	db.Using("default")

	user := User{Id:1}

	err := db.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("no rows")
	} else if err == orm.ErrMissPK {
		fmt.Println("no primary key")
	} else {
		fmt.Println(user.Id,user.Name)
	}

}
