package routers

import (
	"github.com/astaxie/beego"
	"malatang/controllers"
)

func init() {
	beego.Router("/t", &controllers.MainController{})
	beego.Router("/", &controllers.WechatController{})
}
