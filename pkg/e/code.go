package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorAuthCheckTokenFail    = 30001 //token err
	ErrorAuthCheckTokenTimeout = 30002 //token outdated
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001
)
