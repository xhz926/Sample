package main

import (
	"github.com/astaxie/beego"
	_ "hxyhBankTest/models/config"
	_ "hxyhBankTest/routers"
)

func main() {
	beego.Run()
}
