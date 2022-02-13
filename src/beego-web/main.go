package main

import (
	_ "beego-web/routers"

	"github.com/astaxie/beego"
)

func main() {
	// StaticDir["/static"] = "static"

	beego.SetStaticPath("/down1", "download1")
	beego.Run()
}
