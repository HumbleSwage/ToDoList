package api

import (
	"ToDoList/pkg/ctl"
	"ToDoList/pkg/utils"
	"ToDoList/service"
	"ToDoList/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.CreateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.RespError(err))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.ListTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.RespError(err))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}

		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func GetTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.GetTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.GetTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.RespError(err)) // TODO:这个地方的错误处理很奇怪
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.UpdateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.RespError(err))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.RespError(err))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.SearchTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ctl.RespError(err))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			utils.LogrusObj.Info(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
