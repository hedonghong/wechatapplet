package encrypteddata

import (
	"fmt"
	"testing"
	"wechatapplet/common"
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
