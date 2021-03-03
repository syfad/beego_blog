package admin

import (
	"beego_blog/models"
	"github.com/astaxie/beego/orm"
	"strconv"
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
	//fmt.Println(list)
	c.display()
}

// 添加
func (c *LinkController) Add() {
	if c.Ctx.Request.Method == "POST"{
		sitename := c.GetString("sitename")
		url := c.GetString("url")
		rank := c.GetString("rank")
		var link models.Link
		link.Sitename = sitename
		link.Url = url
		//转int类型
		atoi, _ := strconv.Atoi(rank)
		link.Rank = atoi
		if err := link.Insert(); err != nil{
			c.showmsg(err.Error())
		}
		c.Redirect("/admin/link/list", 301)

	}
	c.display()
}

// 编辑
func (c *LinkController) Edit() {
	getInt, _ := c.GetInt("id")
	//这里不太懂，和上面的 var link models.Link,这个用来回显
	link := &models.Link{Id: getInt}
	if err := link.Read(); err != nil{
		c.showmsg("id链接不存在")
	}
	if c.Ctx.Request.Method == "POST"{
		sitename := c.GetString("sitename")
		url := c.GetString("url")
		rank, err := c.GetInt("rank")
		if err != nil{
			rank=0
		}
		link.Sitename = sitename
		link.Url = url
		link.Rank = rank
		if err = link.Update(); err != nil{
			c.showmsg("更新失败")
		}
		c.Redirect("/admin/link/list", 302)
	}
	c.Data["link"] = link
	c.display()
}

//// 删除
//func (c *LinkController) Delete() {
//
//}
