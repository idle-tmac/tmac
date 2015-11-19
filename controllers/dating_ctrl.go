package controllers

import (
	"github.com/astaxie/beego"
        "encoding/json"
        "fmt"
)

type DatingController struct {
	beego.Controller
}

func (c *DatingController) Get() {
        
        msg := c.Ctx.Input.Param(":msg")
        id := c.Ctx.Input.Param(":id")
        idInfo := map[string]string{} 
        idInfo[id] = msg
        b, err := json.Marshal(idInfo)
        if err != nil {
                fmt.Println("error", err)
        }
        c.Ctx.WriteString(string(b))
}
