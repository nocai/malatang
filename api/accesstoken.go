package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessTokenResult struct {
	ApiError

	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken() (string, error) {
	log.Println(AccessTokenURL)

	resp, err := http.Get(AccessTokenURL + "jkjk;")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(body))

	atr := AccessTokenResult{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		return "", err
	}
	atr.checkError()

	return atr.AccessToken, nil
}
