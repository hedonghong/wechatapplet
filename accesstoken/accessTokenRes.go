package accesstoken

import "wechatapplet/common"

type AccessTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	common.Res
}
