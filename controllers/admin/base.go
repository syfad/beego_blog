package admin

import (
	"beego_blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	// 分页
	pager *models.Pager
	// layout.html中需要显示用户名字
	username string
	// 根据html文件命名规则，抽取通用跳转
	// 需要controller名字和方法名
	controllerName string
	methodName     string
	// 加一个标识字段，标识权限验证是否成功
	userid int
}

// 1.获取controller名字和请求方法名，拼接要跳转的页面
// 2.分页相关，创建分页对象等
// 3.权限控制
func (c *BaseController) Prepare() {
	// 获取控制器名和方法名
	// AccountController    Login
	controllerName, actionName := c.GetControllerAndAction()
	// 截取前半段，转小写
	c.controllerName = strings.ToLower(controllerName[:len(controllerName)-10])
	// 方法名转小写
	c.methodName = strings.ToLower(actionName)

	// 校验是否登录了，如果没有登录，应该跳转到登录的请求
	// 进行校验
	c.auth()

	// 分页相关
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	// 定义每页的大小
	pagesize := 5
	// 创建分页对象
	c.pager = models.NewPager(page, pagesize, 0, "")
}

// 权限校验
// 前台请求和登录请求是不需要权限校验的
func (c *BaseController) auth() {
	if c.controllerName == "account" && c.methodName == "login" {
		return
	}
	// 校验cookie是否正确
	// 根据key获取cookie，进行切分
	arr := strings.Split(c.Ctx.GetCookie("auth"), "|")
	// 校验取的数据是否正确
	// 第一个是id，第二个是密码
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		// id转int
		id, _ := strconv.Atoi(idstr)
		if id > 0 {
			// 构造user对象，查询数据库
			user := new(models.User)
			user.Id = id
			// 验证成功
			if user.Read() == nil && user.Password == password {
				c.userid = user.Id
				c.username = user.Username
			}
		}
	}
	// 说明没有登录
	if c.userid == 0 {
		// 跳转到登录请求
		c.Redirect("admin/login", 302)
	}
}

// 实现通用的页面跳转
func (c *BaseController) display(tplname ...string) {
	theme := "admin/"
	// 设置layout不动的部分
	c.Layout = theme + "layout.html"
	c.Data["version"] = beego.AppConfig.String("version")
	c.Data["adminname"] = c.username
	if len(tplname) == 1 {
		c.TplName = theme + tplname[0] + ".html"
	} else {
		// 跳转到具体的tplName页面
		c.TplName = theme + c.controllerName + "_" + c.methodName + ".html"
	}
}

// 公共的异常处理
func (c *BaseController) showmsg(msg ...string) {
	// 设置报错信息：设置成：出错了！
	if len(msg) == 0 {
		msg = append(msg, "出错了！")
	}
	// 设置异常后跳转 redirect
	// 设置报错后，返回上一个页面
	// Referer()可取上一次地址，也可以防盗链
	msg = append(msg, c.Ctx.Request.Referer())
	c.Data["msg"] = msg[0]
	c.Data["redirect"] = msg[1]
	c.display("showmsg")
	// 渲染页面
	err := c.Render()
	if err != nil {
		fmt.Println(err)
	}
	// 终止该次访问
	c.StopRun()
}
