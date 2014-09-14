package controllers

import (
       "fmt"
       "github.com/astaxie/beego"
       "mycomputer/models"
)


type ApiController struct {
	beego.Controller
}

func (this *ApiController) GetUserAll() {
  o := &models.User{}
  fmt.Println(o.Name)
  
  username := this.GetString(":user")
  this.Ctx.Output.Body([]byte(username))

  user := models.GetUser(username)
  this.Ctx.Output.Body([]byte(user.Email))
}

