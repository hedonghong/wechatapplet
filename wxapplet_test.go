package wechatapplet

import (
	"testing"
	"wechatapplet/common"
)

func TestWxapplet_NewClient(t *testing.T) {
	config := &common.AppletConfig{
		Appid:     "",
		AppSecret: "",
		AppHost:   "",
	}
	client := NewClient(config)
	client = client
}
