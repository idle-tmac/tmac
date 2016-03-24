package controllers

import (
	"github.com/astaxie/beego"
        "tmac/lib"
        "fmt"
)

type DongtaiController struct {
	beego.Controller
}

func (c *DongtaiController) Get() {
        
	module := c.Ctx.Input.Param(":module")
        reqWay := c.Ctx.Input.Param(":way")
        content := c.Ctx.Input.Param(":content")
	moduleReqWay := fmt.Sprintf("%s_%s", module, reqWay)
        lib.Call(dongtaiFuncs, moduleReqWay, content, c)

}
func (this *DongtaiController) Post() {
    module := this.GetString("module")
    way := this.GetString("way")
    ticket := this.GetString("ticket")
    fmt.Println(module);
    fmt.Println(way);
    fmt.Println(ticket);
	/*pk := models.GetCruPkg(pkgname)
    if pk.Id == 0 {
        var pp models.PkgEntity
        pp.Pid = 0
        pp.Pathname = pkgname
        pp.Intro = pkgname
        models.InsertPkg(pp)
        pk = models.GetCruPkg(pkgname)
    }
    var at models.Article
    at.Pkgid = pk.Id
    at.Content = content
    models.InsertArticle(at)
    this.Ctx.Redirect(302, "/admin/index")*/
}
