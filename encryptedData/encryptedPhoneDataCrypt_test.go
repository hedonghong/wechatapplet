package encryptedData

import (
	"fmt"
	"testing"
)

func TestEncryptedPhoneDataCrypt_DecryptData(t *testing.T) {
	sessionKey := "5KnrPuuu3KrKn7blXOgYmg=="
	phone, _ := WxBizPhoneDataCrypt.DecryptData(
		sessionKey,
		"9iGbxLwHLXDi/XV27Wt3xM9P8ShbFpwOQReM4G0eqoqZ+0hdZMReRTCrWYc0BujeqXvrhckENoqwt2VI+QNSxAXvh3hFkWp7z5uyB5tNKLmBkZRGko0iQqmMChdFOZX9P5DCGUlQgXmb2dwAVpDLD7OiwL7x336L+V6C9QF4N1vVJASBNV2FnS8c4yL33+aFIwoGTsDeRNKUL9GT5SviVQ==",
		"g2VT9ikWPvaS9uKtj6SDYQ==",
	)
	fmt.Printf("%+v", phone)
	// {"phoneNumber":"18244970044","purePhoneNumber":"18244970044","countryCode":"86","watermark":{"timestamp":1649608924,"appid":"wx361f540e0588c8fd"}}
}
