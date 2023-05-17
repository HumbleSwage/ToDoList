package e

var MsgFlags = map[int]string{
	Success:       "成功",
	Error:         "失败",
	InvalidParams: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
