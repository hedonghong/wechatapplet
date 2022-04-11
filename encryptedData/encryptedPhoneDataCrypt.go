package encryptedData

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"wechatapplet/common"
	"wechatapplet/responses"
)

var WxBizPhoneDataCrypt *wxBizPhoneDataCrypt

type wxBizPhoneDataCrypt struct {
}

func (w *wxBizPhoneDataCrypt) DecryptData(sessionKey, encryptedData, iv string) (phoneInfo *responses.PhoneInfo, err error) {
	// utf8.RuneCountInString()
	if len(sessionKey) != 24 {
		err = &common.CommonError{common.ILLEGALAESKEY, "session_key长度错误"}
		return
	}

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)

	if err != nil {
		return
	}

	if len(iv) != 24 {
		err = &common.CommonError{common.ILLEGALIV, "iv长度错误"}
		return
	}

	aesIv, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		return
	}

	aesCipher, err := base64.StdEncoding.DecodeString(encryptedData)

	if err != nil {
		return
	}

	cipherBlock, err := aes.NewCipher(aesKey)

	if err != nil {
		return
	}

	cipher.NewCBCDecrypter(cipherBlock, aesIv).CryptBlocks(aesCipher, aesCipher)

	aesCipher = PKCS5UnPadding(aesCipher)

	if err = json.Unmarshal(aesCipher, &phoneInfo); err != nil {
		return
	}
	return
}
