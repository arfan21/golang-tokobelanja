package repositorycategory

import (
	"log"

	"github.com/arfan21/golang-tokobelanja/entity"
	"gorm.io/gorm"
)

type RepositoryCategory interface {
	Create(category entity.Category) (entity.Category, error)
	GetAll() ([]entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	Delete(ID uint) error
	IsCategoryExist(ID uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryCategory {
	return &repository{db: db}
}

func (r *repository) Create(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (r *repository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Preload("Products").Find(&categories).Error
	if err != nil {
		return []entity.Category{}, err
	}
	log.Println(categories)
	return categories, nil
}

func (r *repository) Update(category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", category.ID).Updates(&category).First(&category).Error
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (r *repository) Delete(ID uint) error {
	category := entity.Category{}
	category.ID = ID
	err := r.db.First(&category).Where("id = ?", category.ID).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) IsCategoryExist(ID uint) error {
	var category entity.Category
	err := r.db.First(&category, ID).Error
	if err != nil {
		return err
	}

	return nil
}
