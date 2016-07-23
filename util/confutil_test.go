package util

import (
	"fmt"
	"testing"
)

func TestGetAppid(t *testing.T) {
	appid := GetAppid()

	if len(appid) == 0 {
		t.Error("appid is nil")
	}
	fmt.Println("appid", appid)
}

func TestGetSecret(t *testing.T) {
	secret := GetSecret()
	if len(secret) == 0 {
		t.Error("secret is nil")
	}
	fmt.Println("secret", secret)
}

func TestGetToken(t *testing.T) {
	token := GetToken()
	if len(token) == 0 {
		t.Error("token is nil")
	}
	fmt.Println("token", token)
}
