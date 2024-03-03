package handler

import (
	"go-practice/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	u *usecase.AuthUseCase
}

func NewAuthHandler(u *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{u}
}

// TODO: ここにsession認証処理を作る cleanarchitecture的にあっているかは後ほど検討
func (ah *AuthHandler) Login(c *gin.Context) {

}
