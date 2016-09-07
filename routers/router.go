package routers

import (
	"github.com/astaxie/beego"
	"github.com/yunkaiyueming/MonShark/controllers"
)

func init() {

	//数据管理模块
	beego.Router("/", &controllers.MainController{})
	beego.Router("home/", &controllers.HomeController{}, "*:Index")
	beego.Router("home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("home/ShowMgoData", &controllers.HomeController{}, "GET:ShowMgoData")
	beego.Router("home/mgo", &controllers.HomeController{}, "*:GetMongoInfo")

	//用户模块
	beego.Router("user/register", &controllers.UserController{}, "*:Register") //如果这个地方用POST，会导致在控制器中用this.GetString()方法无法获取到数据
	beego.Router("user/doregister", &controllers.UserController{}, "*:DoRegister")
	beego.Router("user/login", &controllers.UserController{}, "*:DoLogin")
	beego.Router("user/logout", &controllers.UserController{}, "*:LogOut")
}
