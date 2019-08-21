package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "home.html"
	//c.Ctx.WriteString("appname:"+beego.AppConfig.String("appname")+
	//	"\nhttpport:"+beego.AppConfig.String("httpport")+
	//	"\nrunmode:"+beego.AppConfig.String("runmode"))
	c.Data["Website"] = "golang.org"
	c.Data["Email"] = "louis@wangke.co"
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	type u struct {
		Name string
		Age  int
		Sex  string
	}
	user := &u{
		"joe",
		10,
		"male",
	}
	c.Data["User"] = user
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	c.Data["Nums"] = nums
	c.Data["TplVar"] ="hey guys"

	c.Data["Html"] = "<div>hello beego</div>"
	c.Data["Pipe"] = "<div>hello beego 2</div>"
}
