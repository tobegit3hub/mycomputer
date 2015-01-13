package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tobegit3hub/mycomputer/models"
)

/* Implement the RESTful server to return JSON */
type ApiController struct {
	beego.Controller
}

/* Receive parameters from form and add user */
func (this *ApiController) AddUser() {
	name := this.GetString("name")
	email := this.GetString("email")
	description := this.GetString("description")

	user := models.User{Name: name, Email: email, Description: description}
	orm.NewOrm().Insert(&user)
}

/* Receive parameter from url and return user object */
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

/* Return user objects */
func (this *ApiController) GetUsers() {
	users := models.GetUsers()

	if users == nil {
		this.Data["json"] = ""
	} else {
		this.Data["json"] = users
	}

	this.ServeJson()
}

/* Receive parameters from form and add item */
func (this *ApiController) AddItem() {
	username := this.GetString("username")
	number, _ := this.GetInt("number")
	image := this.GetString("image")
	description := this.GetString("description")

	item := models.Item{Username: username, Number: number, Image: image, Description: description}
	orm.NewOrm().Insert(&item)
}

/* Receive parameter from url and return user item objects */
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

/* Receive parameters from form and add comment */
func (this *ApiController) AddComment() {
	content := this.GetString("content")

	comment := models.Comment{Content: content}
	orm.NewOrm().Insert(&comment)
}

/* Return the last 9 comment objects */
func (this *ApiController) GetComments() {
	comments := models.GetComments()
	if comments == nil || len(comments) == 0 {
		this.Data["json"] = ""
	} else {
		this.Data["json"] = comments
	}

	this.ServeJson()
}
