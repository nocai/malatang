package api

import (
	"testing"
)

func TestGetCallbackIp(t *testing.T) {
	ipList, err := GetCallbackIp()

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ipList)

}
