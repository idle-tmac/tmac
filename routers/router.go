package routers

import (
	"tmac/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dating/:way&:msg", &controllers.DatingController{})
}
