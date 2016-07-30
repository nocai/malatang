package api

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	accessToken, err := GetAccessToken()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(accessToken)
}
