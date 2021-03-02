package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BaseController
}

// 用户列表
func (c *UserController) List() {
	var list []*models.User
	query := orm.NewOrm().QueryTable(new(models.User))
	_, _ = query.OrderBy("-id").All(&list)

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

// 添加用户
//func (c *UserController) Add() {
//
//}
//
//// 编辑用户
//func (c *UserController) Edit() {
//
//}
//
//// 删除用户
//func (c *UserController) Delete() {
//
//}
