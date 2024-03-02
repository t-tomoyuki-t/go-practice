package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-practice/handler"
	"go-practice/infrastructure/repository"
	"go-practice/router"
	"go-practice/usecase"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_CHARSET"),
		os.Getenv("DB_PARSE_TIME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
