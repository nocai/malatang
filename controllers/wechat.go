package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
	"strings"
)

type WechatController struct {
	beego.Controller
}

func (this *WechatController) Get() {
	signature := this.GetString("signature", "")
	beego.Debug("signature = " + signature)
	timestamp := this.GetString("timestamp", "")
	beego.Debug("timestamp = " + timestamp)
	nonce := this.GetString("nonce", "")
	beego.Debug("nonce = " + nonce)
	echostr := this.GetString("echostr", "")
	beego.Debug("echostr = " + echostr)

	mySign := makeSignature(timestamp, nonce)
	if mySign == signature {
		beego.Info("微信接入成功")
		this.Ctx.WriteString(echostr)
	} else {
		beego.Error("微信接入失败mySign = " + mySign)
		this.Ctx.WriteString("aaaaaaaaaa")
	}
}

func makeSignature(timestamp, nonce string) string {
	var token string = beego.AppConfig.String("token")
	beego.Debug("token = " + token)

	temps := []string{token, timestamp, nonce}
	sort.Strings(temps)

	str := strings.Join(temps, "")
	return str2sha1(str)
}

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
