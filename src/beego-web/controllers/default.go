package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Ctx.WriteString("hello") // 直接输出文本
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// // 不过不指定模板，默认会去到模板目录的 views/MainController/<方法名>.tpl 查找
	c.TplName = "index.tpl"
	// fmt.Print(c)
}

func (u *UserController) Get() {
	u.Data["username"] = "你大爷"
	u.Data["property"] = "快乐"
	u.TplName = "user.html"
}
