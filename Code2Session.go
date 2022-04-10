package wechatapplet

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"wechatapplet/requests"
	"wechatapplet/responses"
)

var Code2Session *code2Session

type code2Session struct {
}

func init() {
	Code2Session = &code2Session{}
}

func (c *code2Session) GetSession(wxLogin requests.WxLogin) (code2Session *responses.Code2SessionData, err error) {
	param := fasthttp.Args{}
	param.Add("appid", Config.Appid)
	param.Add("secret", Config.AppSecret)
	param.Add("js_code", wxLogin.Code)
	param.Add("grant_type", "authorization_code")
	_, bodyByte, err := fasthttp.Get(nil, Config.AuthCode2SessionUri+param.String())
	if err != nil {
		return
	}
	if len(bodyByte) > 0 {
		if err = json.Unmarshal(bodyByte, &code2Session); err != nil {
			return
		}
	}
	return
}
