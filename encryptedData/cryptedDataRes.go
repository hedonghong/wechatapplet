package encrypteddata

import "github.com/hedonghong/wechatapplet/common"

type CryptedDataRes struct {
	OpenId    string            `json:"openId"`
	NickName  string            `json:"nickName"`
	Gender    int8              `json:"gender"`
	Language  string            `json:"language"`
	City      string            `json:"city"`
	Province  string            `json:"province"`
	Country   string            `json:"country"`
	AvatarUrl string            `json:"avatarUrl"`
	UnionId   string            `json:"unionId"`
	Watermark *common.Watermark `json:"watermark"`
}
