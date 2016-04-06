package controllers

import (
	"github.com/astaxie/beego"
        "tmac/lib"
        "fmt"
)

type DatingController struct {
	beego.Controller
}

func (c *DatingController) Get() {
        msg := c.Ctx.Input.Param(":msg")
        way := c.Ctx.Input.Param(":way")
        flag := c.Ctx.Input.Param(":flag")
	fmt.Println("----\n" + flag)
        lib.Call(datingFuncs, way, msg, flag, c)
}
