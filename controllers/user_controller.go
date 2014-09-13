package controllers

import (
       "fmt"
       "github.com/astaxie/beego"
       "mycomputer/models"
)


type UserController struct {
	beego.Controller
}

func (this *UserController) List() {
  o := &models.User{}
  fmt.Println(o.Name)
  this.Ctx.Output.Body([]byte("shorturl"))
}

