package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CallbackIpResult struct {
	ApiError
	IpList []string `json:"ip_list"`
}

func GetCallbackIp() ([]string, error) {
	accessToken, err := GetAccessToken()
	if err != nil {
		return []string{}, err

	}
	callbackIpUrl := strings.Replace(BaseCallbackIpUrl, "ACCESS_TOKEN", accessToken, 1)
	log.Println(callbackIpUrl)

	resp, err := http.Get(callbackIpUrl)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	log.Println(string(body))

	res := CallbackIpResult{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return []string{}, err
	}
	res.checkError()

	return res.IpList, nil
}
