package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Username"] = "kishan.ngm@gmail.com"
	c.Data["Password"] = "Passsword23$%^"
	c.TplName = "index.tpl"
}
