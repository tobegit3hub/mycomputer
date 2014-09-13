package controllers

import (
       "github.com/astaxie/beego"
       "github.com/tobegit3hub/mycomputer/models"
)


type UserController struct {
	beego.Controller
}

/*
func (this *CatalogController) DoAdd() {
	o, err := this.extractCatalog(true)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	_, err = catalog.Save(o)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Redirect("/", 302)
}
*/