package util

import (
	"testing"
)

func TestGetAppid(t *testing.T) {
	appid := GetAppid()

	if len(appid) == 0 {
		t.Error("appid is nil")
	}
	t.Logf("appid", appid)
}

func TestGetSecret(t *testing.T) {
	secret := GetSecret()
	if len(secret) == 0 {
		t.Error("secret is nil")
	}
	t.Logf("secret", secret)
}

func TestGetToken(t *testing.T) {
	token := GetToken()
	if len(token) == 0 {
		t.Error("token is nil")
	}
	t.Logf("token", token)
}
