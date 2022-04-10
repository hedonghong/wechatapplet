package responses

type PhoneInfo struct {
	PhoneNumber     string     `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string     `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string     `json:"countryCode"`     // 区号
	Watermark       *Watermark `json:"watermark"`
}

type PhoneNumberData struct {
	PhoneInfo *PhoneInfo `json:"phone_info"`
	Errcode   int        `json:"errcode"`
	Errmsg    string     `json:"errmsg"`
}
