package main

import (
	"net/http"
	"time"

	"github.com/BKH7/go-client/fetch"
	"github.com/sirupsen/logrus"
)

func main() {
	err := fetch.PostMsg(&http.Client{Timeout: 5 * time.Second}, &fetch.MsgStruct{
		ID:     1,
		Sender: "Tom",
		Msg:    "Hello",
	})
	if err != nil {
		logrus.Error(err)
	}
}
