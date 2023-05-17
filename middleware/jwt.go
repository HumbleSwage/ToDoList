package middleware

import (
	"ToDoList/pkg/ctl"
	"ToDoList/pkg/e"
	"ToDoList/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.Success
		// 检查参数
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少token",
			})
			ctx.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.Error
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(),
			&ctl.UserInfo{Id: claims.Id, UserName: claims.UserName}))

		ctx.Next()
	}
}
