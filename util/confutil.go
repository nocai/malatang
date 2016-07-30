package util

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

const (
	section string = "wechat"
)

var conf *goconfig.ConfigFile

func init() {
	gopath := os.Getenv("GOPATH")
	path := gopath + "/src/malatang/conf/wechat.ini"

	var err error
	conf, err = goconfig.LoadConfigFile(path)
	if err != nil {
		log.Panicln(err)
	}

}

func GetAppid() string {
	return conf.MustValue(section, "appid")

}
func GetSecret() string {
	return conf.MustValue(section, "secret")
}

func GetToken() string {
	return conf.MustValue(section, "token")
}
