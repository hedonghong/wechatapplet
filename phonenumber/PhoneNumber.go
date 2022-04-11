package phonenumber

import (
	"encoding/json"
	"wechatapplet/common"

	"github.com/valyala/fasthttp"
)

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html

const urlPath = "/wxa/business/getuserphonenumber"

type PhoneNumber struct {
	config *common.AppletConfig
	ApiUrl string
}

func NewPhoneNumer(config *common.AppletConfig) *PhoneNumber {
	return &PhoneNumber{
		config: config,
		ApiUrl: config.AppHost + urlPath,
	}
}

func (p *PhoneNumber) GetPhoneNumber(req PhoneNumberReq) (res *PhoneNumberRes, err error) {
	args := &fasthttp.Args{}
	args.Add("access_token", req.AccessToken)
	args.Add("code", req.Code)
	_, bodyByte, err := fasthttp.Post(nil, p.ApiUrl, args)
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
