package api

import (
	"log"
	"malatang/util"
	"strings"
)

const (
	// AccessToken Url
	BaseAccessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
	// CallbackIp Url
	BaseCallbackIpUrl string = "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN"
)

var AccessTokenURL string

func init() {
	appid := util.GetAppid()
	secret := util.GetSecret()
	AccessTokenURL = strings.Replace(strings.Replace(BaseAccessTokenURL, "APPID", appid, 1), "APPSECRET", secret, 1)
	log.Println(AccessTokenURL)
}
