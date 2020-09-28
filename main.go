package main

import (
	"github.com/astaxie/beego"
	"hellobeego0604/db_mysql"
	_ "hellobeego0604/routers"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	//1、连接数据库
	db_mysql.Connect()
	//启动程序
	beego.Run()
}

