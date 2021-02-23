package main

import (
	_ "beego_blog/routers"
	_ "beego_blog/models"
	"github.com/astaxie/beego"
)


func main() {
	beego.Run()
}

