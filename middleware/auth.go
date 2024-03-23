package middleware

import (
	"github.com/gin-gonic/gin"
	"go-practice/domain/repository"
	"net/http"
)

type AuthMiddleware struct {
	ar repository.IAuthRepository
}

func NewAuthMiddleware(ar repository.IAuthRepository) *AuthMiddleware {
	return &AuthMiddleware{ar}
}

func (a *AuthMiddleware) IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, _ := ctx.Cookie("session")
		val, err := a.ar.GetSession(token)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			ctx.Abort()
		}
		if val == "" {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		}
		ctx.Next()
	}
}
