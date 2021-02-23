package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"

	//"strings"
	"time"
)

type Post struct {
	Id       int
	Userid   int
	Author   string
	Title    string
	Color    string
	Content  string `orm:"type(text)"`
	Tags     string
	Views    int
	Status   int
	Posttime time.Time `orm:"type(datetime)"`
	Updated  time.Time `orm:"type(datetime)"`
	Istop    int
	Cover    string
}

func (post *Post) TableName() string {
	//从配置文件读取，表前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "post"
}

//插入
func (post *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(post); err != nil{
		return err
	}
	return nil
}

//删除
func (post *Post) Delete() error {
	if _, err := orm.NewOrm().Delete(post); err != nil{
		return err
	}
	return nil
}

//更新
func (post *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(post,fields...); err != nil{
		return err
	}
	return nil
}

//读取
func (post *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(post,fields...); err != nil{
		return err
	}
	return nil
}

//Link,点击显示文章详情，是一个链接
func (post *Post) Link() string {
	// /article/id
	return "/article/"+strconv.Itoa(post.Id)
}

//ColorTitle,带颜色的标题
func (post *Post) ColorTitle() string {
	if post.Color != ""{
		return fmt.Sprintf("<span style='color:%s'>%s</span>", post.Color, post.Title)
	}
	return post.Title
}

//Excerpt, 文章内容，表里有这字段，以字符串返回即可
func (post *Post) Excerpt() string{
	return post.Content
}

//TagsLink, 文章对应的标签
func (post *Post) TagsLink() string {
	if post.Tags == ""{
		return ""
	}
	//根据,切割
	tagslink := strings.Trim(post.Tags, ",")
	return tagslink
}

//保留删除方法，后面用的时候重写，多表