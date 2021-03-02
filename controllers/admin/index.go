package admin

import (
	"beego_blog/models"
	"github.com/astaxie/beego/orm"
	"os"
	"runtime"
)

type IndexController struct {
	BaseController
}

func (c *IndexController)Index()  {
	//主机名
	c.Data["hostname"],_ = os.Hostname()
	//go版本
	c.Data["gover"] = runtime.Version()
	//操作系统
	c.Data["os"]= runtime.GOOS
	//系统位数
	c.Data["arch"] = runtime.GOARCH
	//文章数量
	c.Data["postnum"],_ = orm.NewOrm().QueryTable(new(models.Post)).Count()
	//分类数量
	c.Data["tagnum"],_ = orm.NewOrm().QueryTable(new(models.Tag)).Count()
	//用户数量
	c.Data["usernum"],_ = orm.NewOrm().QueryTable(new(models.User)).Count()
	//跳转页面
	c.display()
}