package routers

import (
	"github.com/astaxie/beego"
	"whois/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
