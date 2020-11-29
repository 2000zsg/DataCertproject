package routers

import (
	"DataCertproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/lndex.html",&controllers.Index{})
    beego.Router("/lndex", &controllers.Index{})
    beego.Router("/land.html",&controllers.Land{})
    beego.Router("/land", &controllers.Land{})
	beego.Router("/file", &controllers.File{})
    beego.Router("/upload_file.html",&controllers.File{})
    beego.Router("/cert_detail.html",&controllers.CerDetailController{})
	beego.Router("/login_sms.html",&controllers.SmsLoginController{})
    beego.Router("/user_kyc",&controllers.UserKycController{})
	beego.Router("/user_kyc.html",&controllers.UserKycController{})
}
