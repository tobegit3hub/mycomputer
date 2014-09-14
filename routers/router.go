package routers

import (
	"mycomputer/controllers"
	"github.com/astaxie/beego"
)

func init() {

    /* Angular page and router */
    beego.Router("/", &controllers.MainController{})
    beego.Router("/:user", &controllers.MainController{})

    /* Api for angular requests */
    beego.Router("/api/:user", &controllers.ApiController{}, "get:GetUserAll")

}
