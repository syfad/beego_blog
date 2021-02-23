package routers

import (
	"beego_blog/controllers"
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

}
