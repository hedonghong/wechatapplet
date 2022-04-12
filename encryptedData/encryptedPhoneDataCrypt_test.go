package encrypteddata

import (
	"fmt"
	"github.com/hedonghong/wechatapplet/common"
	"testing"
)

func TestEncryptedPhoneDataCrypt_DecryptData(t *testing.T) {
	service := NewEncryptedPhoneDataCrypt(&common.AppletConfig{
		Appid:     "",
		AppSecret: "",
		AppHost:   "",
	})
	sessionKey := "xxx"
	phone, _ := service.DecryptData(
		sessionKey,
		"",
		"",
	)
	fmt.Printf("%+v", phone)
}
