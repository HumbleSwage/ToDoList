package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "request para err",

	ErrorAuthCheckTokenFail:    "Token authentication failed",
	ErrorAuthCheckTokenTimeout: "Token has timed out",
	ErrorAuthToken:             "Token generation failed",
	ErrorAuth:                  "Token error",
	ErrorNotCompare:            "Mismatch",
	ErrorDatabase:              "Database operation error, please try again",
}

// GetMsg Get the corresponding information of the status code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
