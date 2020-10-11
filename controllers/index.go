package controllers

import (
	"DataCertproject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Index struct {
	beego.Controller
}
func (I *Index) Post(){
	var user  models.User
 	err:=I.ParseForm(&user)
	if err!=nil {
		I.Ctx.WriteString("抱歉，解析错误请重试!")
		return
	}
	  _,err =user.SaveUser()
	if err!=nil {
		fmt.Println(err)
		I.Ctx.WriteString("抱歉，用户保存失败，请重试！")
		return
	}
	  I.TplName = "land.html"
}