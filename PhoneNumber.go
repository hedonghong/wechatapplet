package wechatapplet

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"wechatapplet/requests"
	"wechatapplet/responses"
)

var PhoneNumber phoneNumber

type phoneNumber struct {
}

func (p *phoneNumber) GetPhoneNumber(param requests.PhoneNumber) (phoneNumberData *responses.PhoneNumberData, err error) {
	args := &fasthttp.Args{}
	args.Add("access_token", param.AccessToken)
	args.Add("code", param.Code)
	_, bodyByte, err := fasthttp.Post(nil, Config.PhonenumberGetPhoneNumberUri, args)
	if err != nil {
		return
	}
	if len(bodyByte) > 0 {
		if err = json.Unmarshal(bodyByte, &phoneNumberData); err != nil {
			return
		}
	}
	return
}
