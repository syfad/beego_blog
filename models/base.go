package models

import (
	"fmt"
	//加载驱动
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//dhhost = 10.145.102.95
//dbport = 31306
//dbuser = root
//dbpassword = YTMruaGSAoAM6g1f
//dbname = beego_blog
func init() {
	dbhost := beego.AppConfig.String("dhhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	//拼接数据库连接
	//"root:password@tcp(localhost:3306)/beego_blog?charset=utf8"
	dburl := dbuser+":"+ dbpassword +"@tcp("+ dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println(dburl)
	//设置default数据库
	err := orm.RegisterDataBase("default", "mysql", dburl, 50, 20)
	if err != nil{
		fmt.Println("mysql连接失败", err)
	}
	//注册model
	orm.RegisterModel(new(Link),new(Mood),new(Post),new(Tag),new(TagPost),new(User))
	//设置打印sql，日志中打印
	orm.Debug = true
}

//查询最新的4篇文章
func GetLastBlog() []*Post {
	query := orm.NewOrm().QueryTable(&Post{}).Filter("status", 0)
	var result []*Post
	count, _ := query.Count()
	if count>0{
		_, _ = query.OrderBy("-posttime").Limit(4).All(result)

	}
	return result
}

func GetHotBlog() []*Post {
	query := orm.NewOrm().QueryTable(&Post{}).Filter("status", 0)
	var result []*Post
	count, _ := query.Count()
	if count>0{
		_, _ = query.OrderBy("-views").Limit(4).All(result)

	}
	return result
}

//func GetLinks() []*Link {
//	query := orm.NewOrm().QueryTable(&Link{})
//	count, _ := query.Count()
//	if count>0{
//		_, _ = query.OrderBy("-views").Limit(4).All(result)
//
//	}
//	return result
//}