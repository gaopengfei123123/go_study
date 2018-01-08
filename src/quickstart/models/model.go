package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:123123@tcp(127.0.0.1:3306)/go?charset=utf8", 30)
}