package util

import (
	"fmt"
	"github.com/widuu/goini"
)

const (
	section string = "wechat"
	path    string = "G:\\mygo\\src\\malatang\\conf\\wechat.ini"
)

var conf *goini.Config

func init() {
	fmt.Println("path = " + path)
	conf = goini.SetConfig(path)
}

func GetAppid() string {
	fmt.Println(conf.ReadList())
	return conf.GetValue(section, "appid")
}

func GetSecret() string {
	return conf.GetValue(section, "secret")
}

func GetToken() string {
	return conf.GetValue(section, "token")
}
