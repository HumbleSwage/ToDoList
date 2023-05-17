package types

type UserRegisterReq struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
}

type UserLoginReq struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
}

type UserLoginResp struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	CreateAt int64  `json:"create_at"`
}
