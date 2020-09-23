package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//获取get类型的请求参数
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
//以admin 18 为正确数据
if name != "admin"|| age != "18"{
c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
	return
}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
	/*c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1454050489@qq.com"
	c.TplName = "index.tpl"

	 */
}
/**
该post方法是处理post类型的请求的时候要调用的方法
 */
func (c *MainController) Post(){
	fmt.Println("post类型的请求")
	 user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是：",psd)

	//与固定值比较
	if user != "ztc"||psd != "011025"{
		//失败页面
		//c.Ctx.ResponseWriter.Write([]byte("对不起，请求失败"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据准确"))

//request 请求，response响应
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1454050489@qq.com"
	c.TplName = "index.tpl"
}

//}
