package encryptedData

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"wechatapplet/common"
	"wechatapplet/responses"
)

var WXBizDataCrypt *wXBizDataCrypt

type wXBizDataCrypt struct {
	Appid      string
	SessionKey string
}

func init() {
	WXBizDataCrypt = &wXBizDataCrypt{}
}

func (w *wXBizDataCrypt) SetParam(appid, sessionKey string) *wXBizDataCrypt {
	w.Appid = appid
	w.SessionKey = sessionKey
	return w
}

func (w *wXBizDataCrypt) DecryptData(encryptedData, iv string) (cryptedData *responses.CryptedData, err error) {
	// utf8.RuneCountInString()
	if len(w.SessionKey) != 24 {
		err = &common.CommonError{common.ILLEGALAESKEY, "session_key长度错误"}
		return
	}

	aesKey, err := base64.StdEncoding.DecodeString(w.SessionKey)

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

	if err = json.Unmarshal(aesCipher, &cryptedData); err != nil {
		return
	}
	return
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if length-unpadding < 0 {
		return []byte("")
	}
	return src[:(length - unpadding)]
}
