package requests

// https://developers.weixin.qq.com/miniprogram/dev/api/open-api/user-info/wx.getUserProfile.html

type UserInfo struct {
	NickName  string `json:"nickName"`
	Gender    int8   `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

type UserProfile struct {
	ErrMsg        string   `json:"errMsg"`
	RawData       string   `json:"rawData"`
	UserInfo      UserInfo `json:"userInfo"`
	Signature     string   `json:"signature"`
	EncryptedData string   `json:"encryptedData"`
	Iv            string   `json:"iv"`
}
