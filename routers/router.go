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
    //beego.Router("/dongtai/ces/:module?loadTag=3&ticket=0", &controllers.DongtaiController{}, "get:Test")
    beego.Router("/dongtai/cell/:module&:imagename", &controllers.DongtaiController{}, "get:ReqRecommendCell")
    beego.Router("/dongtai/cells/:module&:flag&:num&:ticket", &controllers.DongtaiController{}, "get:ReqRecommendCells")


    //league route
    beego.Router("/league/inschool/cells/:schoolid&:num&:ticket", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCells")
}
