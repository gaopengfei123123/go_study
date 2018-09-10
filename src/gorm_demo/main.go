package main

import (
	"github.com/jinzhu/gorm"
	// 引入 gorm 的 mysql 支持
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 文档地址: https://jasperxu.github.io/gorm-zh/

var db *gorm.DB

// User 的表结构
type User struct {
	ID           int    `gorm:"primary_key"`
	OpenID       string `gorm:"type:varchar(32);not null;index:idx_open"`
	SyUID        int    `gorm:"column:u_id;type:bigint(20);not null"`
	PasswordSalt string `gorm:"type:varchar(32);not null"`
	Info         Info   `gorm:"foreignkey:ID;AssociationForeignKey:ID"`
	Tmps         []Tmp  `gorm:"foreignkey:UID"`
}

// Info 表结构
type Info struct {
	ID       int    `gorm:"primary_key"`
	Avatar   string `gorm:"type:varchar(255);not null"`
	Mobile   string `gorm:"type:varchar(25);not null"`
	Nickname string `gorm:"type:varchar(100);not null"`
}

// Tmp 临时构建的一个表
type Tmp struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(20);not null;"`
	UID  int    `gorm:"column:u_id"`
}

var (
	HOST     string = "127.0.0.1"
	PORT     int    = 33060
	USER     string = "root"
	PASSWORD string = "123123"
	DBNAME   string = "db_open"
)

type Demo struct {
	ID        int    `gorm:"primary_key"`
	IP        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	CreatedAt time.Time
}

func main() {
	// conf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)
	conf := "root:123123@tcp(127.0.0.1:33060)/db_open?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", conf)
	defer db.Close()

	if err != nil {
		fmt.Println("has err ", err)
	}

	// if !db.HasTable(&Demo{}) {
	// 	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Demo{}).Error; err != nil {
	// 		panic(err)
	// 	}
	// }

	// // 新增数据
	// user := &User{
	// 	OpenID: "233asdfasdfads",
	// }
	// if err := db.Create(user).Error; err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("新增数据")
	// }

	// // 查找数据
	// var u User
	// err = db.Model(&User{}).Where("id = ? ", "164").Find(&u).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(u)

	// // 关联查询, 附带条件和 select 查询
	// var u User
	// db.Debug().Select("id, open_id").Where("id = ?", "164").First(&u)
	// db.Debug().Model(&u).Select("id, avatar").Related(&u.Info, "ID")
	// fmt.Println(u)

	// var u User
	// db.Debug().First(&u, 164)
	// db.Debug().Model(&u).Related(&u.Tmps, "UID")
	// fmt.Println(u)

	// 使用关联模式
	var u User
	db.Debug().First(&u, 164)
	// 这里的关键点就在于关联的 Model 需要是一个实体, 有记录的 struct, 然后才能查找对应的数值
	db.Debug().Model(&u).Association("Tmps").Find(&u.Tmps)
	db.Debug().Model(&u).Association("Info").Find(&u.Info)
	fmt.Println(u)

	// var users []Users
	// db.Debug().Preload("Info", "id = ?", "id").Limit(3).Find(&users)
	// fmt.Println(users)

	fmt.Println(db.DB().Ping())

}

// TableName 指定了这个 struct 依赖的表名
func (u User) TableName() string {
	return "tb_u_user"
}

func (i Info) TableName() string {
	return "tb_u_info"
}

func (i Tmp) TableName() string {
	return "tb_tmp"
}
