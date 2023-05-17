package router

import (
	"ToDoList/api"
	"ToDoList/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("create_task", api.CreateTaskHandler())
			authed.POST("list_task", api.ListTaskHandler())
			authed.POST("get_task", api.GetTaskHandler())
			authed.POST("update_task", api.UpdateTaskHandler())
			authed.DELETE("delete_task", api.DeleteTaskHandler())
			authed.POST("search_task", api.SearchTaskHandler())
		}
	}
	return ginRouter
}
