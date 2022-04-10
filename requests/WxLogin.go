package requests

// https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.login.html

type WxLogin struct {
	Code string `json:"code"`
}
