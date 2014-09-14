package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mycomputer/models"
	//       "encoding/json"
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
	o := &models.User{}
	fmt.Println(o.Name)

	username := this.GetString(":user")

	items := models.GetUserItems(username)
	if items == nil || len(items) == 0 {
		this.Data["json"] = ""
	} else {
		this.Data["json"] = items
	}

	this.ServeJson()
}
