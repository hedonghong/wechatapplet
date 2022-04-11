package code2session

import "wechatapplet/common"

type Code2SessionRes struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	common.Res
}
