package api

import (
	"ToDoList/pkg/utils"
	"ToDoList/service"
	"ToDoList/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserSrv()
			resp, err := l.UserRegister(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserSrv()
			resp, err := l.UserLogin(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
