package repository

import (
	"go-practice/domain/entity"
	"go-practice/domain/repository"

	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) repository.IImageRepository {
	return &ImageRepository{db}
}

func (ir *ImageRepository) Get(id int) (*entity.Image, error) {
	image := entity.Image{}
	res := ir.db.First(&image, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &image, nil
}

func (ir *ImageRepository) Save(*entity.Image) (*entity.Image, error) {
	// TODO: Implement
	return nil, nil
}

func (ir *ImageRepository) Delete(*entity.Image) (error) {
	// TODO: Implement
	return nil
}
