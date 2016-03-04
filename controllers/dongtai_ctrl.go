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
