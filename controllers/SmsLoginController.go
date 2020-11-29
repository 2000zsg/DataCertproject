package controllers

import "github.com/astaxie/beego"

type SmsLoginController struct {
	beego.Controller
}

func (s *SmsLoginController)Get()  {
	s.TplName="login_sms.html"
}