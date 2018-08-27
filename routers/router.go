package routers

import (
	"github.com/astaxie/beego"
	"hxyhBankTest/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	for _, v := range controllers.ControllerList {
		beego.Router(v.RootPath, v.ControllerInterface.(beego.ControllerInterface), v.Methods)
	}
}
