package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TagPost struct {
	Id     int
	Tagid  int
	Postid int
	//文章状态
	Poststatus int
	Posttime   time.Time `orm:"type(datetime)"`
}

func (tagpost *TagPost) TableName() string {
	//从配置文件读取，表前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "tag_post"
}

//插入
func (tagpost *TagPost) Insert() error {
	if _, err := orm.NewOrm().Insert(tagpost); err != nil{
		return err
	}
	return nil
}

//删除
func (tagpost *TagPost) Delete() error {
	if _, err := orm.NewOrm().Delete(tagpost); err != nil{
		return err
	}
	return nil
}

//更新
func (tagpost *TagPost) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(tagpost,fields...); err != nil{
		return err
	}
	return nil
}

//读取
func (tagpost *TagPost) Read(fields ...string) error {
	if err := orm.NewOrm().Read(tagpost,fields...); err != nil{
		return err
	}
	return nil
}