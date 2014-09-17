package routers

import (
	"github.com/astaxie/beego"
	"github.com/yvasiyarov/beego_gorelic"
	"mycomputer/controllers"
)

/* Initialize the router */
func init() {
	/* Add NewRelic monitor */
	beego_gorelic.InitNewrelicAgent()

	/* Angular page and router */
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:user", &controllers.MainController{})

	/* Api for angular requests */
	beego.Router("/api/:user", &controllers.ApiController{}, "get:GetUser")
	beego.Router("/api/:user/items", &controllers.ApiController{}, "get:GetUserItems")
	beego.Router("/api/item", &controllers.ApiController{}, "post:AddItem")
}
