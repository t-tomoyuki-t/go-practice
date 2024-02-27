package router

import (
	"go-practice/handler"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, h *handler.TodoHandler) {
	v1 := r.Group("/v1")
	{
		todo := v1.Group("/todos")
		{
			todo.GET("/", func(c *gin.Context) {
				h.GetTodoList(c)
			})
			todo.GET("/:id", func(c *gin.Context) {
				h.GetTodo(c)
			})
			todo.POST("/", func(c *gin.Context) {
				h.Store(c)
			})
			todo.PATCH("/:id", func(c *gin.Context) {
				h.Update(c)
			})
			todo.DELETE("/:id", func(c *gin.Context) {
				h.Delete(c)
			})
		}
	}
}