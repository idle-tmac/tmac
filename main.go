package main

import (
	_ "tmac/routers"
	"tmac/Initial"
	"github.com/astaxie/beego"
)

func main() {
	Initial.Init()
	beego.Run()
}

