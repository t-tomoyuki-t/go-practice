package repository

import (
	"go-practice/domain/entity"
	"go-practice/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) Get(id int) (*entity.User, error) {
	u := entity.User{}
	res := r.db.First(&u, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	u := entity.User{}
	res := r.db.Where("email = ?", email).First(&u)
	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}
