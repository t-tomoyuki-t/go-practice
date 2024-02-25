package main

import (
	"go-practice/handler"
	"go-practice/infrastructure/repository"
	"go-practice/router"
	"go-practice/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := repository.NewTodoReposiory()
	u := usecase.NewTodoUseCase(r)
	h := handler.NewTodoHandler(u)
	e := gin.Default()
	router.SetRoutes(e, h)
	e.Run(":8080")
}
