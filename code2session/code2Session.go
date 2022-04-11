package code2session

import (
	"encoding/json"
	"wechatapplet/common"

	"github.com/valyala/fasthttp"
)

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

const urlPath = "/sns/jscode2session"

type Code2Session struct {
	config *common.AppletConfig
	ApiUrl string
}

func NewCode2Session(config *common.AppletConfig) *Code2Session {
	return &Code2Session{
		config: config,
		ApiUrl: config.AppHost + urlPath,
	}
}

func (c *Code2Session) GetSession(req Code2SessionReq) (res *Code2SessionRes, err error) {
	param := fasthttp.Args{}
	param.Add("appid", c.config.Appid)
	param.Add("secret", c.config.AppSecret)
	param.Add("js_code", req.Code)
	param.Add("grant_type", "authorization_code")
	_, bodyByte, err := fasthttp.Get(nil, c.ApiUrl+"?"+param.String())
	if err != nil {
		return
	}
	if len(bodyByte) > 0 {
		if err = json.Unmarshal(bodyByte, &res); err != nil {
			return
		}
	}
	return
}
