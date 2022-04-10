package wechatapplet

import (
	"fmt"
	"testing"
)

var (
	//appid         = "wx4f4bc4dec97d474b"
	//sessionKey    = "tiihtNczf5v6AKRyjwEUhQ=="
	//encryptedData = `CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew==`
	//iv            = "r7BXXKkLb8qrSNn05n0qiA=="

	appid         = "wx361f540e0588c8fd"
	sessionKey    = "5KnrPuuu3KrKn7blXOgYmg=="
	encryptedData = `gfk1nQahszJprqUuLNT6S5i6uQv7puIlLchipOZDp4TXr/ZKtoNM+F01TG63mzYt54QTOvd3qbYKO2+Cvl0cXdkY78ldK/WACGtzrlZqPgJfjxPphKok0VtHvQ3s9t24Ow1Gx+UlMqCRM5X7eyytiWY3MfZcZn/y+qtyiDi3UAaRsdf4xBKdtNwszkfsnk846CxXTvmRUu00HpBjqUbm0cPXaJ5XyHAowndFgby/UPRYa6uWTkQRgggJQrMC7vM6yAIXZbllESTEY40wO3LbkH9tP8Wa8qMweE4hLrIw98jVwysaLNDuNaLxNAfVs4Z32t6e204ugo9jB9o3EiAXXjuGqlvhp61YyjxUIUQQB5hmlSTxX5cxgDqhSFECk09tXefTj/ikiM34sPH91NVMRDNo1dmEXDXUONMyy9bZZYA=`
	iv            = "TyC6mfELXlWnjn0TUy6pPg=="
)

func TestWXBizDataCrypt_DecryptData(t *testing.T) {
	object, _ := WXBizDataCrypt.SetParam(appid, sessionKey).DecryptData(encryptedData, iv)
	fmt.Printf("%+v", object)
}
