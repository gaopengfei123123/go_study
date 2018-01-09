package models
import (
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type user struct {
	Id   int	
	Name string `orm:"size(100)"`
	Password string
}

// UserModal 用户表模型
type UserModal struct {
    user
	IsError bool
	Error string
}


func init(){
	orm.RegisterModel(new(user))
}

// GetOne 获取一条用户信息
func (th *UserModal) GetOne() (*UserModal){
	db := orm.NewOrm()
	db.Using("default")

	u := user{Id:th.Id}
	err := db.Read(&u)
	if err == orm.ErrNoRows {
		th.Error = "no rows"
		th.IsError = true
	} else if err == orm.ErrMissPK {
		th.IsError = true
	} 
	th.user = u

	return th
}

/*
Login 用户登录
@username
@password
*/
func (th *UserModal) Login(username string,password string) (*UserModal){

	u := user{Id:th.Id,Name:th.Name}
	th.user = u
	return th
}
