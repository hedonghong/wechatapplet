package code2session

import (
	"fmt"
	"github.com/hedonghong/wechatapplet/common"
	"testing"
)

func TestCode2Session_GetSession(t *testing.T) {
	service := NewCode2Session(&common.AppletConfig{
		Appid:     "",
		AppSecret: "",
		AppHost:   "",
	})
	session, _ := service.GetSession(Code2SessionReq{
		Code: "021Uvp0w3Xi6jY2bQ63w34WIZj4Uvp0z",
	})
	fmt.Printf("%+v", session)
}
