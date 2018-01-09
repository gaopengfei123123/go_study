package models
import (
	"encoding/hex"
	"crypto/sha256"
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

const CONNECTION = "default"
// User 表结构
type User struct {
	Id   int	
	Name string `orm:"size(100)"`
	Password string
	Token string `orm:"size(64)"`
}

// UserModal 用户表模型
type UserModal struct {
    User
	IsError bool
	Error string
}


func init(){
	orm.RegisterModel(new(User))
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


	return th
}

func (th *UserModal) checkPassword(password string, user User) (bool) {
	return password == user.Password
}

func (th *UserModal) generateHash() string {
	var str string = th.Name + th.Password
	var data []byte = []byte(str)
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func (th *UserModal) flashToken() bool {
	db := orm.NewOrm()
	db.Using(CONNECTION)
	return false
}
