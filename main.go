package main

import (
	"DataCertproject/qkl_mysql"
	_ "DataCertproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	qkl_mysql.Qkl()
	beego.Run()
}

