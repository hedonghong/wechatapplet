package accesstoken

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"wechatapplet/common"
	"wechatapplet/responses"
)

var AccessToken accessToken

type accessToken struct {
}

func (g *accessToken) GetAccessToken() (accessToken *responses.AccessTokenData, err error) {
	param := fasthttp.Args{}
	param.Add("appid", common.Config.Appid)
	param.Add("secret", common.Config.AppSecret)
	param.Add("grant_type", "client_credential")
	_, bodyByte, err := fasthttp.Get(nil, common.Config.AuthGetAccessTokenUri+param.String())
	if err != nil {
		return
	}
	if len(bodyByte) > 0 {
		if err = json.Unmarshal(bodyByte, &accessToken); err != nil {
			return
		}
	}
	return
}
