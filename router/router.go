package router

import (
	"go-practice/handler"
	"go-practice/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, th *handler.TodoHandler, ah *handler.AuthHandler, uh *handler.UserHandler, am *middleware.AuthMiddleware) {
	v1 := r.Group("/v1")
	{
		v1.POST("register", func(c *gin.Context) {
			uh.RegisterUser(c)
		})
		v1.POST("login", func(c *gin.Context) {
			ah.Login(c)
		})

		todo := v1.Group("/todos")
		todo.Use(am.IsAuthenticated())
		{
			todo.GET("/", func(c *gin.Context) {
				th.GetTodoList(c)
			})
			todo.GET("/:id", func(c *gin.Context) {
				th.GetTodo(c)
			})
			todo.POST("/", func(c *gin.Context) {
				th.Store(c)
			})
			todo.PATCH("/:id", func(c *gin.Context) {
				th.Update(c)
			})
			todo.DELETE("/:id", func(c *gin.Context) {
				th.Delete(c)
			})
		}

		user := v1.Group("/users")
		{
			user.GET("/:id", func(c *gin.Context) {
				uh.GetUser(c)
			})
		}
	}
}
