package controllers

import (
	"github.com/astaxie/beego"
        "tmac/lib"
        _"fmt"
)

type DatingController struct {
	beego.Controller
}

func (c *DatingController) Get() {
        msg := c.Ctx.Input.Param(":msg")
        way := c.Ctx.Input.Param(":way")
        lib.Call(datingFuncs, way, msg, c)
}
