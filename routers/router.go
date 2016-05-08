package routers

import (
	"tmac/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dating/:way&:msg&:flag", &controllers.DatingController{})
    beego.Router("/dating/:way&:msg", &controllers.DatingController{})

    //login and register
    beego.Router("/login/:phone&:passwd", &controllers.LoginController{}, "get:LoginCheck")


    //dongtai route
    beego.Router("/dongtai/cell/:module&:imagename", &controllers.DongtaiController{}, "get:ReqRecommendCell")
    beego.Router("/dongtai/cells/:module&:flag&:num&:ticket", &controllers.DongtaiController{}, "get:ReqRecommendCells")


    //league route
    beego.Router("/league/inschool/cells/:schoolid&:num&:ticket", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCells")
    beego.Router("/league/inschool/cell/:schoolid&:leagueid", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCell")
   
}
