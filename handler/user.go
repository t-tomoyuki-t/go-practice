package handler

import (
	"go-practice/domain/entity"
	"go-practice/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	u *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{u}
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uh.u.Get(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	newUser, err := uh.u.Register(&user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"user": newUser})
}
