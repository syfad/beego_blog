package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"

	"github.com/astaxie/beego/validation"
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
func (c *UserController) Add() {
	// 用于报错回显的map
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if c.Ctx.Request.Method == "POST"{
		username := c.GetString("username")
		password := c.GetString("password")
		password2 := c.GetString("password2")
		email := c.GetString("email")
		active, _ := c.GetInt("active")
		valid := validation.Validation{}
		input["username"] = username
		input["password"] = password
		input["password2"] = password2
		input["email"] = email
		//判断用户名非空，长度
		if result := valid.Required(username, "username"); !result.Ok{
			errmsg["username"] = "用户不能为空"
		}else if result := valid.MaxSize(username,14,"username"); !result.Ok{
			errmsg["username"] = "用户名密码长度不能大于15个字符！"
		}
		// 判断密码非空
		if result := valid.Required(password, "password"); !result.Ok {
			// 处理异常
			errmsg["password"] = "密码不能为空!"
		}
		// 2次密码是否一致
		if result := valid.Required(password2, "password2"); !result.Ok {
			// 处理异常
			errmsg["password2"] = "确认密码不能为空!"
		} else if password != password2 {
			errmsg["password2"] = "两次输入的密码不一致!"
		}
		// 邮箱校验
		if result := valid.Required(email, "email"); !result.Ok {
			// 处理异常
			errmsg["email"] = "邮箱不能为空!"
		} else if result := valid.Email(email, "email"); !result.Ok {
			errmsg["email"] = "邮箱不合法!"
		}
		// 严谨处理
		if active > 0 {
			active = 1
		} else {
			active = 0
		}
		//添加操作
		if len(errmsg) ==0{
			//添加用户入库
			var user = &models.User{}
			//组装数据库
			user.Username = username
			user.Password = Md5([]byte(password))
			user.Email = email
			user.Active = active
			if err := user.Insert(); err != nil{
				c.showmsg(err.Error())
			}
			//跳转用户列表
			c.Redirect("/admin/user/list",302)
		}
	}
	c.Data["errmsg"] = errmsg
	c.Data["input"] = input
	c.display()
}

// 编辑用户
func (c *UserController) Edit() {
	id, _ := c.GetInt("id")
	user := &models.User{Id: id}
	if err := user.Read(); err != nil{
		c.showmsg("用户不存在")
	}
	errmsg := make(map[string]string)
	if c.Ctx.Request.Method == "POST"{
		password := strings.TrimSpace(c.GetString("password"))
		password2 := strings.TrimSpace(c.GetString("password2"))
		email := strings.TrimSpace(c.GetString("email"))
		active, _ := c.GetInt("active")

		//申明validation对象
		valid := validation.Validation{}
		// 判断密码非空
		if result := valid.Required(password, "password"); !result.Ok {
			// 处理异常
			errmsg["password"] = "密码不能为空!"
		} else if result := valid.Required(password2, "password2"); !result.Ok {
			errmsg["password2"] = "确认密码不能为空!"
		} else if password != password2 {
			errmsg["password2"] = "两次输入的密码不一致!"
		} else {
			user.Password = Md5([]byte(password))
		}
		// 邮箱校验
		if result := valid.Required(email, "email"); !result.Ok {
			// 处理异常
			errmsg["email"] = "邮箱不能为空!"
		} else if result := valid.Email(email, "email"); !result.Ok {
			errmsg["email"] = "邮箱不合法!"
		} else {
			user.Email = email
		}
		// 严谨处理
		if active > 0 {
			user.Active = 1
		} else {
			user.Active = 0
		}
		//更新入库
		if len(errmsg) == 0{
			err := user.Update()
			if err != nil{
				c.showmsg(err.Error())
			}
			c.Redirect("/admin/user/list", 302)
		}
	}
	c.Data["user"] = user
	c.Data["errmsg"] = errmsg
	c.display()
}

// 删除用户
func (c *UserController) Delete() {
	id, _ := c.GetInt("id")
	user := &models.User{Id: id}
	if err := user.Read(); err == nil{
		if err := user.Delete(); err != nil{
			c.showmsg("删除失败")
		}
		c.Redirect("/admin/user/list", 302)
	}
}
