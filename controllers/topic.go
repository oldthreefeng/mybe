/**
 * @time: 2019-08-22 23:46
 * @author: louis
 */
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"mybee/models"
)

type TopicController struct {
	beego.Controller
}

func (tc *TopicController) Get()   {
	tc.TplName = "topic.html"
	tc.Data["IsTopic"] = true
	tc.Data["IsLogin"] = checkAccount(tc.Ctx)  //将状态判断为是否登录
}

func (tc *TopicController) Post()  {
	if !checkAccount(tc.Ctx) {
		tc.Redirect("/login",302)
		return
	}
	title := tc.Input().Get("title")
	content := tc.Input().Get("content")
	var err error
	err = models.AddTopic(title,content)
	if err != nil {
		logs.Error(err)
	}
	tc.Redirect("/topic",302)
}

func (tc *TopicController) Add() {
	tc.TplName = "topic_add.html"
}