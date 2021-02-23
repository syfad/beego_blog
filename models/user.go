package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Username string
	Password string
	Email string
	Logincount int
	//加密后的password，cookie会使用
	Authkey string
	Active int
}

func (user *User) TableName() string {
	//从配置文件读取，表前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "user"
}

//插入
func (user *User) Insert() error {
	if _, err := orm.NewOrm().Insert(user); err != nil{
		return err
	}
	return nil
}

//删除
func (user *User) Delete() error {
	if _, err := orm.NewOrm().Delete(user); err != nil{
		return err
	}
	return nil
}

//更新
func (user *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(user,fields...); err != nil{
		return err
	}
	return nil
}

//读取
func (user *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(user,fields...); err != nil{
		return err
	}
	return nil
}