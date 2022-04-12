package encrypteddata

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"github.com/hedonghong/wechatapplet/common"
)

// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html

type EncryptedDataCrypt struct {
	config *common.AppletConfig
}

func NewEncryptedDataCrypt(config *common.AppletConfig) *EncryptedDataCrypt {
	return &EncryptedDataCrypt{
		config: config,
	}
}

func (e *EncryptedDataCrypt) DecryptData(sessionKey, encryptedData, iv string) (res *CryptedDataRes, err error) {
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

	aesCipher = common.PKCS5UnPadding(aesCipher)

	if err = json.Unmarshal(aesCipher, &res); err != nil {
		return
	}
	return
}
