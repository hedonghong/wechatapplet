package common

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if length-unpadding < 0 {
		return []byte("")
	}
	return src[:(length - unpadding)]
}
