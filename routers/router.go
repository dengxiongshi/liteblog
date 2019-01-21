package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	//注解路由
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
	beego.ErrorController(&controllers.ErrorController{})
}
