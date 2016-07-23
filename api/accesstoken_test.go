package api

import (
	"fmt"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	accessToken, err := GetAccessToken()
	if err != nil {
		t.Error("accessToken is nil")
	}
	fmt.Println("accessToken = " + accessToken)
}
