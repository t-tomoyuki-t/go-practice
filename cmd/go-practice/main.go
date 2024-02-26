package main

import (
	"go-practice/handler"
	"go-practice/infrastructure/repository"
	"go-practice/router"
	"go-practice/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDb()
	r := repository.NewTodoReposiory(db)
	u := usecase.NewTodoUseCase(r)
	h := handler.NewTodoHandler(u)
	e := gin.Default()
	router.SetRoutes(e, h)
	e.Run(":8080")
}

func initDb() *gorm.DB {
	dsn := "user:password@tcp(mysql:3306)/go-practice?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	return db
}
