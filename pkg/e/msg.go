package e

var msgMapping = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "参数异常",

	ERROR_EXIST_TAG:         "标签已经存在",
	ERROR_NOT_EXIST_TAG:     "标签不存在",
	ERROR_NOT_EXIST_ARTICLE: "文章不存在",

	ERROR_AUTH:                     "认证失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "TOKEN 鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "TOKEN 已过期",
}

func GetMssg(code int) (msg string) {
	msg, ok := msgMapping[code]
	if !ok {
		msg = msgMapping[ERROR]
	}
	return
}
