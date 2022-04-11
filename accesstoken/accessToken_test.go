package accesstoken

import (
	"fmt"
	"testing"
)

func TestAccessToken_GetAccessToken(t *testing.T) {
	token, _ := AccessToken.GetAccessToken()
	fmt.Printf("%+v", token)
	// {"access_token":"55_ofkyDXlzK_veB0yq0mYoxmMLj-6JZSBbsqep5XBgBrblTwSF7VYn7ttumZf-AE09oXrk2-mouA_i2DcU3NXR72H1m7IhHjqYy4dAZC9SEgJTTyFLlJWUTrpKNbBp4bIuW-TXLyHhPeMK3AqdMOFcABAHJA","expires_in":7200}
}
