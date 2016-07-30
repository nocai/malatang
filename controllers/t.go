package controllers

import (
	"github.com/astaxie/beego"
	"malatang/api"
)

type TController struct {
	beego.Controller
}

func (this *TController) Get() {
	at, err := api.GetAccessToken()
	if err == nil {
		this.Ctx.WriteString(at)
	}

}
