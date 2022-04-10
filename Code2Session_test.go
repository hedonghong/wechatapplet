package wechatapplet

import (
	"fmt"
	"testing"
	"wechatapplet/requests"
)

func TestCode2Session_GetSession(t *testing.T) {
	session, _ := Code2Session.GetSession(requests.WxLogin{
		Code: "021Uvp0w3Xi6jY2bQ63w34WIZj4Uvp0z",
	})
	fmt.Printf("%+v", session)
}
