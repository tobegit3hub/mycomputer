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

func (this *ApiController) GetUserAll() {
  o := &models.User{}
  fmt.Println(o.Name)
  
  username := this.GetString(":user")
//  this.Ctx.Output.Body([]byte(username))

/*  user := models.GetUser(username)


  if user == nil {
    this.Ctx.Output.Body([]byte("nil"))
  }else {
  this.Ctx.Output.Body([]byte(user.Email))
  }
*/

items := models.GetUserItems(username)
if items == nil || len(items) == 0{
    this.Ctx.Output.Body([]byte("nil"))
  }else {
//  this.Ctx.Output.Body([]byte("ok"))
  }


this.Data["json"] = items
this.ServeJson()


}

