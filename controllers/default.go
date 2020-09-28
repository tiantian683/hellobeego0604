package controllers

import (
"encoding/json"
"fmt"
"github.com/astaxie/beego"
"hellobeego0604/models"
"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//name1 := c.GetString("name")
	//age1,err := c.GetInt("age")
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
func (c *MainController) Post() {
	//Body := c.Ctx.Request.Body
	dataByes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("数据接收失败重试")
		return
	}
	var person models.Person
	err = json.Unmarshal(dataByes,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名:",person.User,"年龄:",person.Age)
	c.Ctx.WriteString("用户密码是:"+person.User)
}

//}
