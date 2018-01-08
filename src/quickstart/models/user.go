package models
import (
	"fmt"
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)
// Model Struct
type User struct {
    Id   int
	Name string `orm:"size(100)"`
	IsError bool
	Error string
}

func init(){
	orm.RegisterModel(new(User))
}

func (this *User) GetOne() (User){
	db := orm.NewOrm()
	db.Using("default")

	user := User{Id:this.Id}


	err := db.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("no rows")
	} else if err == orm.ErrMissPK {
		fmt.Println("no primary key")
	} else {
		// fmt.Println
	}

	return user
}

/*
Login 用户登录
@username
@password
*/
func (c *User) Login(username string,password string) (User){

	user := User{Id:c.Id,Name:c.Name,IsError:false,Error:""}
	return user
}
