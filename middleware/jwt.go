package middleware

import (
	"ToDoList/pkg/e"
	"ToDoList/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.CheckToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
