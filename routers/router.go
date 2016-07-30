package routers

import (
	"github.com/astaxie/beego"
	"malatang/controllers"
)

func init() {
	beego.Router("/t", &controllers.TController{})
	beego.Router("/", &controllers.WechatController{})
}
