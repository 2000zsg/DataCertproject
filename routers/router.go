package routers

import (
	"DataCertproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/lndex", &controllers.Index{})
    beego.Router("/land", &controllers.Land{})
}
