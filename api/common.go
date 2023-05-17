package api

import (
	"ToDoList/pkg/ctl"
	"ToDoList/pkg/e"
	"encoding/json"
)

func ErrorResponse(err error) *ctl.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespErrorWithData(err, "JSON类型不匹配")
	}
	return ctl.RespErrorWithData(err, "参数错误", e.InvalidParams)
}
