package db_mysql

import (

	"database/sql"
	"fmt"
	"github.com/astaxie/beego"

)

//var Db

func Connect()  {
	//定义config变量，接收并赋值全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称",appName)
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请查阅后重试")
	}
	fmt.Println("应用监听端口",port)

	driver := config.String("db_driver")
	dbuser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_Ip")
	dbName := config.String("db_name")
	//db.Open连接数据库，有俩个参数
	db,err :=sql.Open(driver,dbuser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil{
		//err不为nil，表示数据库连接出现错误，程序就此中断
		//早发现，早解决
		panic("数据连接打开失败，，请重试")//使程序进入恐慌状态，程序终止执行
	}
	//Db = db
	fmt.Println(db)
}
//将用户信息保存到数据库中去
//func AddUser(u models.User) (int64,error) {
 //Db.Exec("inser into user (name")
	//md5Hash := md5.New()
	//md5Hash :=

