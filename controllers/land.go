package controllers

import (
	"DataCertproject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Land struct {
	beego.Controller
}
func (L *Land)Post(){
	var user models.User
	err:= L.ParseForm(&user)
	if err!=nil {
		L.Ctx.WriteString("抱歉，用户信息解析失败,请重试!")
		return
	}
	u, err := user.QueryUser()
	if err!=nil {
		fmt.Println(err)
		L.Ctx.WriteString("抱歉,登入失败，请重试！")
		return
	}
	L.Data["phone"] = u.Phone
	L.TplName = "uoload.html"
}