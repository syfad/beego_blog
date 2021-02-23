package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Mood struct {
	Id       int
	Content  string `orm:"type(text)"`
	//图片路径
	Cover    string `orm:"size(70)"`
	//发表时间
	Posttime time.Time `orm:"type(datetime)"`
}

func (mood *Mood) TableName() string {
	//从配置文件读取，表前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "mood"
}

//插入
func (mood *Mood) Insert() error {
	if _, err := orm.NewOrm().Insert(mood); err != nil{
		return err
	}
	return nil
}

//删除
func (mood *Mood) Delete() error {
	if _, err := orm.NewOrm().Delete(mood); err != nil{
		return err
	}
	return nil
}

//更新
func (mood *Mood) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(mood,fields...); err != nil{
		return err
	}
	return nil
}

//读取
func (mood *Mood) Read(fields ...string) error {
	if err := orm.NewOrm().Read(mood,fields...); err != nil{
		return err
	}
	return nil
}