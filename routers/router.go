package routers

import (
	"beego_blog/controllers"
	"beego_blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	//前台页面
	//首页
    beego.Router("/", &controllers.MainController{},"*:Index")
    //分页路由
    //定义正则路由，从路由可以接受API参数page，指定为int类型
	beego.Router("/index:page:int.html", &controllers.MainController{},"*:Index")

    //后台页面

    //首页
    beego.Router("/admin", &admin.IndexController{}, "*:Index")

    //权限管理
    beego.Router("/admin/login", &admin.AccountController{}, "*:Login")

    //退出登录
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")

    //个人信息
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")

    //碎言碎语
	beego.Router("/admin/mood/list", &admin.MoodController{}, "*:List")
	beego.Router("/admin/mood/add", &admin.MoodController{}, "*:Add")
	beego.Router("/admin/mood/delete", &admin.MoodController{}, "*:Delete")

	//友链
	beego.Router("/admin/link/list", &admin.LinkController{}, "*:List")
	beego.Router("/admin/link/add", &admin.LinkController{}, "*:Add")
	beego.Router("/admin/link/edit", &admin.LinkController{}, "*:Edit")
	//beego.Router("/admin/link/delete", &admin.LinkController{}, "*:Delete")

	//user
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
}
