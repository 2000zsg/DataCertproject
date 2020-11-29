package controllers

import (
	"DataCertproject/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type Land struct {
	beego.Controller
}

func (I *Land) Get() {
	I.TplName = "land.html"
}
func (L *Land) Post() {
	var user models.User
	err := L.ParseForm(&user)
	if err != nil {
		L.Ctx.WriteString("抱歉，用户信息解析失败,请重试!")
		return
	}
	u, err := user.QueryUser()
	//fmt.Println("====",u.Name)
	if err != nil {
		fmt.Println(err)
		L.Ctx.WriteString("抱歉,登入失败，请重试！")
		return
	}
	name:=strings.TrimSpace(u.Name)
	Card:=strings.TrimSpace(u.Card)
	Sex:=strings.TrimSpace(u.Sex)
	if name ==""||Card==""||Sex==""{
		L.Data["Phone"]=u.Phone
		L.TplName="user_kyc.html"
		return
	}
	L.Data["Phone"] = u.Phone
	L.TplName = "uoload.html"
}
