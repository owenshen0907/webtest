package routers

import (
	"webtest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/content", &controllers.ContentController{})

}
