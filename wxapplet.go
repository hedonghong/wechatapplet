package wechatapplet

import (
	"github.com/hedonghong/wechatapplet/accesstoken"
	"github.com/hedonghong/wechatapplet/code2session"
	"github.com/hedonghong/wechatapplet/common"
	"github.com/hedonghong/wechatapplet/encrypteddata"
	"github.com/hedonghong/wechatapplet/phonenumber"
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

func (c *Client) NewEncryptedDataCrypt() *encrypteddata.EncryptedDataCrypt {
	return encrypteddata.NewEncryptedDataCrypt(c.Config)
}

func (c *Client) NewEncryptedPhoneDataCrypt() *encrypteddata.EncryptedPhoneDataCrypt {
	return encrypteddata.NewEncryptedPhoneDataCrypt(c.Config)
}
