package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id int
	Name string
	Count int
}

func (tag *Tag) TableName() string {
	//从配置文件读取，表前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "tag"
}

//插入
func (tag *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(tag); err != nil{
		return err
	}
	return nil
}

//删除
func (tag *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(tag); err != nil{
		return err
	}
	return nil
}

//更新
func (tag *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(tag,fields...); err != nil{
		return err
	}
	return nil
}

//读取
func (tag *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(tag,fields...); err != nil{
		return err
	}
	return nil
}