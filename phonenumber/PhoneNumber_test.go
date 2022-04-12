package phonenumber

import (
	"fmt"
	"github.com/hedonghong/wechatapplet/common"
	"testing"
)

func TestPhoneNumber_GetPhoneNumber(t *testing.T) {
	service := NewPhoneNumer(&common.AppletConfig{
		Appid:     "",
		AppSecret: "",
		AppHost:   "",
	})
	phone, _ := service.GetPhoneNumber(PhoneNumberReq{
		Code:        "",
		AccessToken: "55_ofkyDXlzK_veB0yq0mYoxmMLj-6JZSBbsqep5XBgBrblTwSF7VYn7ttumZf-AE09oXrk2-mouA_i2DcU3NXR72H1m7IhHjqYy4dAZC9SEgJTTyFLlJWUTrpKNbBp4bIuW-TXLyHhPeMK3AqdMOFcABAHJA",
	})
	fmt.Printf("%+v", phone)
}
