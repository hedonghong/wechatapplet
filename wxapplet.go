package wechatapplet

import (
	"wechatapplet/accesstoken"
	"wechatapplet/code2session"
	"wechatapplet/common"
	"wechatapplet/phonenumber"
)

type Client struct {
	Config *common.AppletConfig
}

func NewClient(config *common.AppletConfig) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) NewAccessToken() *accesstoken.AccessToken {
	return accesstoken.NewAccessToken(c.Config)
}

func (c *Client) NewCode2Session() *code2session.Code2Session {
	return code2session.NewCode2Session(c.Config)
}

func (c *Client) NewPhoneNumer() *phonenumber.PhoneNumber {
	return phonenumber.NewPhoneNumer(c.Config)
}
