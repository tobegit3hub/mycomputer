package routers

import (
	"mycomputer/controllers"
	"github.com/astaxie/beego"
)

func init() {

    /* Angular home page */
    beego.Router("/", &controllers.MainController{})

    /* Api for angular requests */
    beego.Router("/api/:user", &controllers.UserController{}, "get:List")

}
