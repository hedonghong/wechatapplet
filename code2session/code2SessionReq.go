package code2session

// https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.login.html

type Code2SessionReq struct {
	Code string `json:"code"`
}
