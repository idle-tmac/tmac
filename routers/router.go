package routers

import (
	"tmac/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dating/:way&:msg&:flag", &controllers.DatingController{})
    beego.Router("/dating/module", &controllers.DatingController{},"get:Test")
//http://localhost:8080/blog/30/beego/true/98.45?page=10

    //register
    beego.Router("/register/verifycode", &controllers.RegisterController{}, "get:RegisterDeal")
    //beego.Router("/register/verifycode?phone=xxx&verifycode=xxx", &controllers.RegisterController{}, "get:RegisterDeal")
    beego.Router("/register/userinfo", &controllers.RegisterController{}, "get:RegisterPasswdUpLoad")
    //beego.Router("/register/userinfo?phone=xxx&passwd1=xxx&passwd2=xxx", &controllers.RegisterController{}, "get:RegisterPasswdUpLoad")
    
    //login
    beego.Router("/login", &controllers.LoginController{}, "get:LoginCheck")
    //beego.Router("/login?loginway=xxx&value=xxx&passwd=xxx", &controllers.LoginController{}, "get:LoginCheck")


    //dongtai route
    beego.Router("/dongtai/cell", &controllers.DongtaiController{}, "get:ReqRecommendCell")
    //beego.Router("/dongtai/cell?module=xxx&imagename=xxx", &controllers.DongtaiController{}, "get:ReqRecommendCell")
    beego.Router("/dongtai/cells", &controllers.DongtaiController{}, "get:ReqRecommendCells")
    //beego.Router("/dongtai/cells?module=xxx&flag=xxx&num=xxx&ticket=xxx", &controllers.DongtaiController{}, "get:ReqRecommendCells")


    //league route
    beego.Router("/league/inschool/cells", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCells")
    //beego.Router("/league/inschool/cells?schoolid=xxx&num=xxx&ticket=xxx", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCells")
    beego.Router("/league/inschool/cell", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCell")
    //beego.Router("/league/inschool/cell?schoolid=xxx&leagueid=xxx", &controllers.LeagueController{}, "get:ReqinSchoolLeagueCell")
   
}
