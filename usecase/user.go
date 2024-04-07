package usecase

import (
	"go-practice/domain/entity"
	"go-practice/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserExtarnal interface {
	SendRegisterd(*entity.User) error
}

type UserUseCase struct {
	r  repository.IUserRepository
	ex IUserExtarnal
}

func NewUserUseCase(r repository.IUserRepository, ex IUserExtarnal) *UserUseCase {
	return &UserUseCase{r, ex}
}

func (u *UserUseCase) Get(id int) (*entity.User, error) {
	user, err := u.r.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usecase *UserUseCase) Register(user *entity.User) (*entity.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), bcrypt.DefaultCost)
	newUser := entity.NewUser(
		user.Id,
		user.Name,
		user.Email,
		string(hash),
	)

	savedUser, err := usecase.r.Save(newUser)
	if err != nil {
		return nil, err
	}

	err = usecase.ex.SendRegisterd(savedUser)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}
