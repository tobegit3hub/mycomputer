package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tobegit3hub/mycomputer/models"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) GetUser() {
	username := this.GetString(":user")
	user := models.GetUser(username)

	if user == nil {
		this.Data["json"] = ""
	} else {
		this.Data["json"] = user
	}

	this.ServeJson()
}

func (this *ApiController) GetUserItems() {
	username := this.GetString(":user")

	items := models.GetUserItems(username)
	if items == nil || len(items) == 0 {
		this.Data["json"] = ""
	} else {
		this.Data["json"] = items
	}

	this.ServeJson()
}

func (this *ApiController) AddItem() {
	//var item models.Item // ORM doesn't work
	//this.ParseForm(&item)
	//orm.NewOrm().Insert(&item)

	username := this.GetString("username")
	number, _ := this.GetInt("number")
	image := this.GetString("image")
	description := this.GetString("description")

	item := models.Item{Username: username, Number: number, Image: image, Description: description}
	orm.NewOrm().Insert(&item)
}
