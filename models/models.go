package models

import (
       "github.com/astaxie/beego/orm"
       _ "github.com/go-sql-driver/mysql"
)

type User struct {
     Name        string `orm:"pk"` 
     Email         string
     Description       string
//    Posts []*Post `orm:"reverse(many)"` 
}

type Item struct {
     Number                    int64 `orm:"pk"` 
     Image                 string
     Description	   string
     Username		   string
     User *User `orm:"rel(fk)"`
}

type Comment struct {
     Id	     int64 `orm:"pk;auto"`
     Content string
}

func (*User) TableEngine() string {
     return engine()
}

func (*Item) TableEngine() string {
     return engine()
}


func engine() string {
     return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
//     orm.RegisterModelWithPrefix("mycomputer_", new(User), new(Item))
     orm.RegisterModel(new(User), new(Item))
     orm.RegisterDriver("mysql", orm.DR_MySQL)
    
    orm.RegisterDataBase("default", "mysql", "root:@/mycomputer?charset=utf8")
}

