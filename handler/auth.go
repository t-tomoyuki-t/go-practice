package handler

import (
	"go-practice/domain/entity"
	"go-practice/domain/vo"
	"go-practice/usecase"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	au *usecase.AuthUseCase
}

func NewAuthHandler(au *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{au}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var a entity.Auth
	err := c.BindJSON(&a)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := ah.au.Login(a)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	secure, _ := strconv.ParseBool(os.Getenv("COOKIE_SECURE"))
	httpOnly, _ := strconv.ParseBool(os.Getenv("COOKIE_HTTP_ONLY"))
	c.SetCookie("session", token.String(), vo.TTL_SECOND, "/", os.Getenv("APP_DOMAIN"), secure, httpOnly)
}
