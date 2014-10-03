package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name        string `orm:"pk"`
	Email       string
	Description string
}

type Item struct {
	Id          int64  `orm:"pk;auto"`
	Username    string //`form:"username"`
	Number      int64  //`form:"number"`
	Image       string //`form:"image"`
	Description string //`form:"description"`
}

type Comment struct {
	Id      int64 `orm:"pk;auto"`
	Content string
}

func (*User) TableEngine() string {
	return engine()
}

func (*Item) TableEngine() string {
	return engine()
}

func (*Comment) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
	orm.RegisterModel(new(User), new(Item), new(Comment))
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	db_user := beego.AppConfig.String("db_user")
	db_password := beego.AppConfig.String("db_password")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_database := beego.AppConfig.String("db_database")

	// orm.RegisterDataBase("mycomputer", "mysql", "root1:561801src1@tcp(104.131.61.2521:33061)/mycomputer")
	orm.RegisterDataBase("default", "mysql", db_user+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_database+"?charset=utf8")
}

func GetUser(username string) *User {
	user := User{Name: username}
	err := orm.NewOrm().Read(&user, "Name")
	if err != nil {
		return nil
	}
	return &user
}

func GetUserItems(username string) []*Item {
	var items []*Item
	_, err := orm.NewOrm().QueryTable("item").Filter("username", username).OrderBy("-number").All(&items)
	if err != nil {
		return nil
	}
	return items
}

func GetComments() []*Comment {
	var comments []*Comment
	_, err := orm.NewOrm().QueryTable("comment").Limit(10).OrderBy("-id").All(&comments)
	if err != nil {
		return nil
	}
	return comments
}
