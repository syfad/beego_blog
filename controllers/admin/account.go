package admin

import (
	"beego_blog/models"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

type AccountController struct {
	BaseController
}

// 登录方法
func (c *AccountController) Login() {
	// 用隐藏域判断请求方式
	if c.GetString("dosubmit") == "yes" {
		// post请求，进行表单提交
		// 需要：account、password、remember
		// 获取3个参数，取掉两边空格
		account := strings.TrimSpace(c.GetString("account"))
		password := strings.TrimSpace(c.GetString("password"))
		remember := strings.TrimSpace(c.GetString("remember"))
		// 校验用户输入是否正确，是否登录通过
		if account != "" && password != "" {
			// 构建user对象
			var user = &models.User{}
			user.Username = account
			// 需要将用户输入的password，进行加密后，与数据库password比对
			// 这里我们用md5实现的数据加密，实际中为了更安全，进行加盐加密
			if user.Read("username") != nil || user.Password != Md5([]byte(password)) {
				// 用户查不到，或者密码比对失败
				c.Data["errmsg"] = "账号或密码错误！"
				// 判断账号是否激活
			} else if user.Active == 0 {
				c.Data["errmsg"] = "账号尚未激活！"
			} else {
				// 登录成功
				// 1.登录次数加1
				// 2.设置cookie，key是auth，value是id|md5(pwd)
				user.Logincount += 1
				// 更新数据库+1
				err := user.Update("logincount")
				if err != nil {
					fmt.Println(err)
				}
				// 设置cookie，判断是否记住一周
				if remember == "yes" {
					// 设置一周的有效期
					c.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+
						Md5([]byte(password)), 60*60*24*7)
				} else {
					// 正常设置cookie，默认有效是3600秒
					c.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+
						Md5([]byte(password)))
				}
				// 重定向到后台首页
				c.Redirect("/admin", 302)
			}
		}
	}
	// get请求，就是跳转到登录页面
	c.TplName = "admin/" + c.controllerName + "_" + c.methodName + ".html"
}

// md5加密
func Md5(buf []byte) string {
	// 创建md5对象
	mymd5 := md5.New()
	// 进行加密
	mymd5.Write(buf)
	// 获取哈希值
	result := mymd5.Sum(nil)
	return fmt.Sprintf("%x", result)
}

// 退出登录
func (c *AccountController) Logout() {
	// 覆盖掉cookie的值即可
	// 若用的session，可以直接调用delete即可
	c.Ctx.SetCookie("auth", "")
	c.Redirect("/admin/login", 302)
}

// 个人信息
func (c *AccountController) Profile() {
	// 构建user，通过userid，可以拿到当前登录的用户
	user := &models.User{Id: c.userid}
	if err := user.Read(); err != nil {
		c.showmsg(err.Error())
	}
	// false不展示：修改成功的提示
	updated := false
	errmsg := make(map[string]string)
	if c.Ctx.Request.Method == "POST" {
		// 用户如果是post请求，则调用提交表单
		password := strings.TrimSpace(c.GetString("password"))
		newpassword := strings.TrimSpace(c.GetString("newpassword"))
		newpassword2 := strings.TrimSpace(c.GetString("newpassword2"))
		// 判断不为空
		if newpassword != "" {
			// 判断原密码是否正确
			if password == "" || Md5([]byte(password)) != user.Password {
				errmsg["password"] = "当前密码错误！"
				// 新密码不能小于6位
			} else if len(newpassword) < 6 {
				errmsg["newpassword"] = "新密码不能小于6位字符！"
			} else if newpassword != newpassword2 {
				errmsg["newpassword2"] = "两次输入的密码不一致！"
			}
		}
		// 没有问题后，密码更新
		if len(errmsg) == 0 {
			// 新密码的md5处理
			user.Password = Md5([]byte(newpassword))
			// 更新数据库
			err := user.Update("password")
			if err != nil {
				fmt.Println(err)
			}
			updated = true
		}
	}
	// 若是get请求，则跳转页面
	c.Data["updated"] = updated
	c.Data["user"] = user
	c.Data["errmsg"] = errmsg
	// 自动跳转
	c.display()
}
