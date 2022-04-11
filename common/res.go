package common

type Res struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}
