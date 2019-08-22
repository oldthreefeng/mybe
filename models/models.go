/**
 * @time: 2019-08-22 00:02
 * @author: louis
 */
package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
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
	TopicTime       time.Time `orm:"index"`
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
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddTopic(title,content string) error {
	o := orm.NewOrm() //获取orm对象
	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
	}
	_, err := o.Insert(topic) // 不存在,则创建
	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm() //获取orm对象
	cate := &Category{
		Created:   time.Now(), //不为nil, 创建表的时候是不为null
		Title:     name,
		TopicTime: time.Now(),
	} //创建category对象
	qs := o.QueryTable("category")            //查询
	err := qs.Filter("title", name).One(cate) //如果cate是slice,用ALl,传递指针
	if err == nil {
		return err //已经存在
	}
	_, err = o.Insert(cate) // 不存在,则创建
	if err != nil {
		return err
	}
	return nil

}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64) //获取string-->int64
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid} //删除
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
