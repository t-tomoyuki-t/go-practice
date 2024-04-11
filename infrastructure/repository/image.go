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

func (ir *ImageRepository) Save(image *entity.Image) (*entity.Image, error) {
	res := ir.db.Save(image)
	if res.Error != nil {
		return nil, res.Error
	}
	return image, nil
}

func (ir *ImageRepository) Delete(image *entity.Image) error {
	res := ir.db.Delete(image)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (ir *ImageRepository) GetPublicImage(image *entity.Image) (*entity.Image, error) {
	return nil, nil
}

func (ir *ImageRepository) Upload(image *entity.Image) (*entity.Image, error) {
	return nil, nil
}
