package wechatapplet

const (
	ILLEGALAESKEY = -41001
	ILLEGALIV     = -41002
	ILLEGALBUFFER = -41003
	DECODEBASE64  = -41004

	SUCCESS     = 0
	SERVICEBUSY = -1

	CODE2SESSION_CODEINVALID = 40029 // code 无效
	CODE2SESSION_LIMIT       = 45011 // 频率限制，每个用户每分钟100次
	CODE2SESSION_RISK        = 40226 // 高风险等级用户，小程序登录拦截

	ACCESSTOKEN_APPSECRET = 40001 // AppSecret 错误或者 AppSecret 不属于这个小程序，请开发者确认 AppSecret 的正确性
	ACCESSTOKEN_GRANTTYPE = 40002 // 请确保 grant_type 字段值为 client_credential
	ACCESSTOKEN_APPID     = 40013 // 不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写

	PHONENUMBER_CODE = 40029 // 不合法的code（code不存在、已过期或者使用过）
)
