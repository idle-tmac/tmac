package routers

import (
	"tmac/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dating/:way&:msg&:flag", &controllers.DatingController{})
    beego.Router("/dating/:way&:msg", &controllers.DatingController{})

    //dongtai route
    beego.Router("/dongtai/cell/:module&:imagename", &controllers.DongtaiController{}, "get:ReqRecommendCell")
    beego.Router("/dongtai/cells/:module&:flag&:num&:ticket", &controllers.DongtaiController{}, "get:ReqRecommendCells")
}
