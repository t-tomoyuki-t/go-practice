package usecase

import "go-practice/domain/repository"

type AuthUseCase struct {
	r repository.IAuthRepository
}

func NewAuthUseCase(r repository.IAuthRepository) *AuthUseCase {
	return &AuthUseCase{r}
}
