/**
 * @time: 2019-08-22 20:50
 * @author: louis
 */
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"mybee/models"
)

type CategoryController struct {
	beego.Controller
}

func (cc *CategoryController) Get() {
	if !checkAccount(cc.Ctx) {  //判断是否登录,没有登录就跳转登录
		cc.Redirect("/login",302)
		return
	}
	op := cc.Input().Get("op")
	switch op {
	case "add":
		name := cc.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			logs.Error(err)
		}
		cc.Redirect("/category", 302)
		return
	case "del":
		id := cc.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			logs.Error(err)
		}
		cc.Redirect("/category", 302)
		return
	}
	cc.Data["IsLogin"] = checkAccount(cc.Ctx)
	cc.Data["IsCategory"] = true
	cc.TplName = "category.html"
	var err error //已经声明的类型,且不能更改的,只能重新生成
	cc.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		logs.Error(err)
	}
}
