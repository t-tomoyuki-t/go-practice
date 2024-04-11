package repository

import "go-practice/domain/entity"

type IImageRepository interface {
	Get(int) (*entity.Image, error)
	Save(*entity.Image) (*entity.Image, error)
	Delete(*entity.Image) (error)
	GetPublicImage(*entity.Image) (*entity.Image, error)
	Upload(*entity.Image) (*entity.Image, error)
}
