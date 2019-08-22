/**
 * @time: 2019-08-22 01:29
 * @author: louis
 */
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) Get() {
	isExit := lc.Input().Get("exit") == "true"
	if isExit {
		lc.Ctx.SetCookie("uname", "", -1, "/") //-1 立即删除cookie然后重定向
		lc.Ctx.SetCookie("pwd", "", -1, "/")
		lc.Redirect("/", 301)
		return
	}
	lc.TplName = "login.html"
}

func (lc *LoginController) Post() {
	//lc.Ctx.WriteString(fmt.Sprint(lc.Input()))

	uname := lc.Input().Get("uname")
	pwd := lc.Input().Get("pwd")
	autoLogin := lc.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		lc.Ctx.SetCookie("uname", uname, maxAge, "/")
		lc.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}//只是将cookie存入了,但是为读出

	lc.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")  //cookie如果共享
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}
