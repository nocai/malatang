package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"malatang/util"
	"net/http"
	"strings"
)

var accessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"

func init() {
	appid := util.GetAppid()
	fmt.Println("appid", appid)
	accessTokenURL = strings.Replace(accessTokenURL, "APPID", appid, 1)
	secret := util.GetSecret()
	accessTokenURL = strings.Replace(accessTokenURL, "APPSECRET", secret, 1)
}

type AccessTokenResult struct {
	Access_token string
	Expires_in   int
}

func GetAccessToken() (string, error) {
	resp, err := http.Get(accessTokenURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))

	atr := AccessTokenResult{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(atr)

	return atr.Access_token, nil
}
