package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mybee/models"
	_ "mybee/routers"
)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug =true  //pro,dev 设置不同真值
	orm.RunSyncdb("default",false,true)
	beego.Run()
}
