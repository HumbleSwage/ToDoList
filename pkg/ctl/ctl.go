package ctl

import "ToDoList/pkg/e"

// Response 基础响应
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 响应列表
type DataList struct {
	Item  interface{} `json:"item"`
	Total int         `json:"total"`
}

// TokenData 用户登录专属带用户信息和token返回
type TokenData struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}

// TrackErrorResponse 带有追踪信息的返回
type TrackErrorResponse struct {
	Response
	TraceId string `json:"trace_id"`
}

// RespSuccess 成功响应
func RespSuccess(code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Msg:    e.GetMsg(status),
	}
}

// RespSuccessWithData 响应成功数据
func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}

// RespError 失败响应
func RespError(err error, code ...int) *Response {
	status := e.Error
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Error:  err.Error(),
	}
}

// RespErrorWithData 失败响应数据
func RespErrorWithData(err error, data interface{}, code ...int) *Response {
	status := e.Error
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   data,
		Error:  err.Error(),
	}
}
