package models
import (
	// "fmt"
	"encoding/hex"
	"crypto/sha256"
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

const CONNECTION = "default"

// Role 表结构
type UserRoles struct {
	Id int
	Role string
	Title string
	Uid int
}

// User 表结构
type User struct {
	Id   int	
	Name string `orm:"size(100)"`
	Password string
	Token string `orm:"size(64)"`
	// UserRoles	[]*UserRoles `orm:"reverse(many)"`
}



// UserModal 用户表模型
type UserModal struct {
    User
	IsError bool
	Error string
	Ip string
}


func init(){
	orm.RegisterModel(new(User),new(UserRoles))
}

// GetOne 获取一条用户信息
func (th *UserModal) GetOne(column string) (*UserModal){
	db := orm.NewOrm()
	db.Using(CONNECTION)

	err := db.Read(&th.User,column)

	if err == orm.ErrNoRows {
		th.Error = "no rows"
		th.IsError = true
	} else if err == orm.ErrMissPK {
		th.IsError = true
	} 

	return th
}

/*
Login 用户登录
@username
@password
*/
func (th *UserModal) Login() (*UserModal){

	pwd := th.Password
	th.GetOne("Name")

	if th.IsError {
		return th
	}

	if !th.checkPassword(pwd, th.User) {
		th.IsError = true
		th.Error = "error password or username"
	}

	th.Token = th.generateHash()
	if !th.flashToken() {
		th.IsError = true
		th.Error = "flash Token error"
	}

	return th
}

func (th *UserModal) checkPassword(password string, user User) (bool) {
	return password == user.Password
}

func (th *UserModal) generateHash() string {
	str := th.Name + th.Password + th.Ip
	data := []byte(str)
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func (th *UserModal) flashToken() bool {
	db := orm.NewOrm()
	db.Using(CONNECTION)

	if _, err := db.Update(&th.User,"Token"); err != nil {
		return false
	}

	return true
}
