package routers

import (
	"tmac/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dating/:way&:msg&:flag", &controllers.DatingController{})
    beego.Router("/dating/:way&:msg", &controllers.DatingController{})

    //register
    beego.Router("/register/verifycode/:phone&:verifycode", &controllers.RegisterController{}, "get:RegisterDeal")
    beego.Router("/register/userinfo/:phone&:passwd1&:passwd2", &controllers.RegisterController{}, "get:RegisterPasswdUpLoad")
    
    //login
    beego.Router("/login/:loginway&:value&:passwd", &controllers.LoginController{}, "get:LoginCheck")


    //dongtai route
    beego.Router("/dongtai/cell/:module&:imagename", &controllers.DongtaiController{}, "get:ReqRecommendCell")
    beego.Router("/dongtai/cells/:module&:flag&:num&:ticket", &controllers.DongtaiController{}, "get:ReqRecommendCells")


    //league route
    beego.Router("/league/inschool/cells/:schoolid&:num&:ticket", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCells")
    beego.Router("/league/inschool/cell/:schoolid&:leagueid", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCell")
   
}
