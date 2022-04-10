package responses

type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type CryptedData struct {
	OpenId    string     `json:"openId"`
	NickName  string     `json:"nickName"`
	Gender    int8       `json:"gender"`
	Language  string     `json:"language"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	Country   string     `json:"country"`
	AvatarUrl string     `json:"avatarUrl"`
	UnionId   string     `json:"unionId"`
	Watermark *Watermark `json:"watermark"`
}
