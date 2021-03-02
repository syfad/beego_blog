package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type LinkController struct {
	BaseController
}

// 列表
func (c *LinkController) List() {
	var list  []*models.Link
	query := orm.NewOrm().QueryTable(new(models.Link))
	_, _ = query.OrderBy("-rank").All(&list)
	//query := all
	//query.All(&list)
	c.Data["list"]=list
	fmt.Println(list)
	c.display()
}

// 添加
func (c *LinkController) Add() {
	if c.Ctx.Request.Method == "POST"{
		sitename := c.GetString("sitename")
		url := c.GetString("url")
		rank := c.GetInt("rank")
		var link models.Link
		link.Sitename = sitename
		link.Url = url
		link.Rank = rank

	}
}

//// 编辑
//func (c *LinkController) Edit() {
//
//}
//
//// 删除
//func (c *LinkController) Delete() {
//
//}
