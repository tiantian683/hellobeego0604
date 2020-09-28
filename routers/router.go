package routers

import (
	"github.com/astaxie/beego"
	"hellobeego0604/controllers"
)

func init() {
//用户注册功能接口 http://127.0.0.1:8080/register
beego.Router("/register",&controllers.RegisterController{})
//http://127.0.0.1:8080
beego.Router("/", &controllers.MainController{})
}

