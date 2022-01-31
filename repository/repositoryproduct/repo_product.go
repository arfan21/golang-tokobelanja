package repositoryproduct

import (
	"github.com/arfan21/golang-tokobelanja/entity"
	"gorm.io/gorm"
)

type RepositoryProduct interface {
	Create(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	DeleteByID(ID uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryProduct {
	return &repository{db: db}
}

func (r *repository) Create(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (r *repository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}

func (r *repository) Update(product entity.Product) (entity.Product, error) {
	err := r.db.Where("id = ?", product.ID).Updates(&product).First(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (r *repository) DeleteByID(ID uint) error {
	product := entity.Product{}
	product.ID = ID
	err := r.db.First(&product).Where("id = ?", product.ID).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
