package phonenumber

import (
	"fmt"
	"testing"
	"wechatapplet/requests"
)

func TestPhoneNumber_GetPhoneNumber(t *testing.T) {
	phone, _ := PhoneNumber.GetPhoneNumber(requests.PhoneNumber{
		Code:        "",
		AccessToken: "55_ofkyDXlzK_veB0yq0mYoxmMLj-6JZSBbsqep5XBgBrblTwSF7VYn7ttumZf-AE09oXrk2-mouA_i2DcU3NXR72H1m7IhHjqYy4dAZC9SEgJTTyFLlJWUTrpKNbBp4bIuW-TXLyHhPeMK3AqdMOFcABAHJA",
	})
	fmt.Printf("%+v", phone)
}
