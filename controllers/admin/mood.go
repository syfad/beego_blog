package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
)

type MoodController struct {
	BaseController
}

// 心情列表，查询的方法
func (c *MoodController) List() {
	// 分页加排序
	var list []*models.Mood
	// 创建orm，得到句柄
	query := orm.NewOrm().QueryTable(new(models.Mood))
	// 查询总数
	count, _ := query.Count()
	c.pager.SetTotalnum(int(count))
	c.pager.SetUrlpath("/admin/mood/list?page=%d")
	// 严谨判断，进行分页查询
	if count > 0 {
		// 设置分页查询偏移量
		// 第1页	从索引0开始查询
		// 第2页	从索引2开始查询
		// 第3页	从索引4开始查询
		// 偏移量 = (当前页码-1) * 每页大小
		offset := (c.pager.Page - 1) * c.pager.Pagesize
		// 根据id降序
		_, err := query.OrderBy("-id").Limit(c.pager.Pagesize, offset).All(&list)
		if err != nil {
			fmt.Println("err = ", err)
		}
	}
	// 返回数据
	c.Data["list"] = list
	c.Data["pagebar"] = c.pager.PageBar()
	// 跳转页面
	c.display()
}

// 添加的方法
func (c *MoodController) Add() {
	// 判断请求类型是否是post
	if c.Ctx.Request.Method == "POST" {
		// 获取内容
		content := c.GetString("content")
		// 构建mood对象
		var mood models.Mood
		mood.Content = content
		// 随机取一个图片，作为心情的配图
		rand.Seed(time.Now().Unix())
		// 随机出0到10的数字，用于拼接取图
		var r = rand.Intn(10)
		mood.Cover = "/static/upload/blog" + fmt.Sprintf("%d", r) + ".jpg"
		// 心情的发表时间
		mood.Posttime = time.Now()
		// 插入数据库
		if err := mood.Insert(); err != nil {
			c.showmsg(err.Error())
		}
		c.Redirect("/admin/mood/list", 302)
	}
	// 页面跳转
	c.display()
}

// 删除的方法
func (c *MoodController) Delete() {
	// 获取id
	id, err := c.GetInt("id")
	if err != nil {
		c.showmsg("删除失败!")
	}
	// 删除
	mood := models.Mood{Id: id}
	if err = mood.Read(); err == nil {
		err := mood.Delete()
		if err != nil {
			c.showmsg("删除失败!")
		}
	}
	c.Redirect("/admin/mood/list", 302)
}
