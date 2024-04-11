package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"go-practice/handler"
	"go-practice/infrastructure/external"
	"go-practice/infrastructure/repository"
	"go-practice/middleware"
	"go-practice/router"
	"go-practice/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := initDb()
	rdb := initInMemory()
	uex := external.NewUserSlackNotification()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUseCase(ur, uex)
	uh := handler.NewUserHandler(uu)

	ar := repository.NewAuthRepository(rdb)
	au := usecase.NewAuthUseCase(ar, ur)
	ah := handler.NewAuthHandler(au)

	tr := repository.NewTodoReposiory(db)
	tu := usecase.NewTodoUseCase(tr)
	th := handler.NewTodoHandler(tu)

	am := middleware.NewAuthMiddleware(ar)

	e := gin.Default()
	router.SetRoutes(e, th, ah, uh, am)
	e.Run(":8080")
}

func initDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
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

func initInMemory() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("INMEMORY_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("INMEMORY_HOST"),
			os.Getenv("INMEMORY_PORT"),
		),
		Password: os.Getenv("INMEMORY_PASSWORD"),
		DB:       db,
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ping result", pong)

	return rdb
}

func initStorage() *s3.Client {
	ctx := context.Background()

	endpoint := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: os.Getenv("AWS_S3_BUCKET"),
		}, nil
	})

	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred), config.WithEndpointResolverWithOptions(endpoint))
	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg, func (options *s3.Options)  {
		options.UsePathStyle = true
	})
}
