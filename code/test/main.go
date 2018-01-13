package main
import (
	"fmt"
	"time"
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

const CONNECTION = "default"

// Role 表结构
type Role struct {
	Id int
	Role string
	Title string
    // Uid int
    // User *User  `orm:"rel(fk)"`
}

// User 表结构
type User struct {
	Id   int	
	Name string `orm:"size(100)"`
	Password string
	Token string `orm:"size(64)"`
	// UserRoles  []*UserRoles `orm:"reverse(many)"`
}



// UserModal 用户表模型
type UserModal struct {
    User
	IsError bool
	Error string
	Ip string
}

func init(){
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:123123@tcp(127.0.0.1:3306)/go?charset=utf8", 30)
	orm.RegisterModel(new(User),new(Role))
}

func main() {
    // user := UserModal{}
    // user.User.Id = 3
    // data := user.GetOne("Id")

	// fmt.Println("export result:",data)
	

	// fmt.Println("goruntime begin")
	// syncChan1 := make(chan struct{},1)
	// syncChan2 := make(chan struct{},2)
	// go receive(strChan,syncChan1,syncChan2)
	// go send(strChan,syncChan1,syncChan2)
	// <-syncChan1
	// <-syncChan2
	// fmt.Println("goruntime end")

	userChan := make(chan *UserModal,2)
	go getMany(userChan,1)
	go getMany(userChan,3)
	for elem := range userChan{
		fmt.Println(elem)
	}

	close(userChan)
	
	
}

func getMany(userChan chan<- *UserModal,id int){
	modal := UserModal{}
	modal.User.Id = id
	data := modal.GetOne("Id")
	time.Sleep(time.Second * 5)
	userChan <- data
}

var strChan = make(chan string,3)
func receive(strChan <-chan string,syncChan1 chan<- struct{},syncChan2 chan<- struct{}) {
	syncChan1<- struct{}{}
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for elem := range strChan {
		fmt.Println("Received:", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send (strChan chan<- string,syncChan1 chan<- struct{},syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a","b","c","d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
		}
	}
	fmt.Println("wait 2 seconds.. [sender]")
	time.Sleep(time.Second * 5)
	close(strChan)
	syncChan2 <- struct{}{}

}



// GetOne 获取一条用户信息
func (th *UserModal) GetOne(column string) (*UserModal){
	db := orm.NewOrm()
    db.Using(CONNECTION)

    err := db.Read(&th.User,column)
   
    fmt.Println("reverse:",err)

	if err == orm.ErrNoRows {
		th.Error = "no rows"
		th.IsError = true
	} else if err == orm.ErrMissPK {
		th.IsError = true
	} 

	return th
}

