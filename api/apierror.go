package api

import (
	"fmt"
	"log"
)

type ApiError struct {
	Code int64  `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (this ApiError) checkError() {
	if this.Code != 0 {
		log.Panicln(this)
	}
}

func (this ApiError) Error() string {
	return fmt.Sprintf("code: %d message: %s", this.Code, this.Msg)
}
