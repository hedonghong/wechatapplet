package accesstoken

import (
	"encoding/json"
	"github.com/hedonghong/wechatapplet/common"

	"github.com/valyala/fasthttp"
)

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html

const urlPath = "/cgi-bin/token"

type AccessToken struct {
	config *common.AppletConfig
	ApiUrl string
}

func NewAccessToken(config *common.AppletConfig) *AccessToken {
	return &AccessToken{
		config: config,
		ApiUrl: config.AppHost + urlPath,
	}
}

func (a *AccessToken) GetAccessToken() (res *AccessTokenRes, err error) {
	param := fasthttp.Args{}
	param.Add("appid", a.config.Appid)
	param.Add("secret", a.config.AppSecret)
	param.Add("grant_type", "client_credential")
	_, bodyByte, err := fasthttp.Get(nil, a.ApiUrl+"?"+param.String())
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
