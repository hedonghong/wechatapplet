package requests

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html

type PhoneNumber struct {
	AccessToken string `json:"accessToken"`
	Code        string `json:"code"`
}
