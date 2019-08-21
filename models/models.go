/**
 * @time: 2019-08-22 00:02
 * @author: louis
 */
package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/mybee.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"` //创建时间
	Views           int64     `orm:"index"` //浏览次数
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"` //创建时间
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"` //浏览次数
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category),new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}
