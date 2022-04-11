package common

var Config *config

type config struct {
	Appid                        string
	AppSecret                    string
	AppHost                      string
	//AuthCode2SessionUri          string
	//AuthGetAccessTokenUri        string
	//PhonenumberGetPhoneNumberUri string
}

func init() {
	Config = &config{
		Appid:                        "wx361f540e0588c8fd",
		AppSecret:                    "cc88a6d94a5b740ad66345af3ceb8282",
		AppHost:                      "https://api.weixin.qq.com",
		//AuthCode2SessionUri:          "https://api.weixin.qq.com/sns/jscode2session?",
		//AuthGetAccessTokenUri:        "https://api.weixin.qq.com/cgi-bin/token?",
		//PhonenumberGetPhoneNumberUri: "https://api.weixin.qq.com/wxa/business/getuserphonenumber?",
	}
}
