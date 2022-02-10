package serviceproduct

import (
	"github.com/arfan21/golang-tokobelanja/entity"
	"github.com/arfan21/golang-tokobelanja/model/modelproduct"
	"github.com/arfan21/golang-tokobelanja/repository/repositorycategory"
	"github.com/arfan21/golang-tokobelanja/repository/repositoryproduct"
	"github.com/arfan21/golang-tokobelanja/validation"
	"github.com/jinzhu/copier"
)

type ServiceProduct interface {
	Create(product modelproduct.Request) (modelproduct.Response, error)
	GetAll() ([]modelproduct.Response, error)
	Update(product modelproduct.Request) (modelproduct.Response, error)
	DeleteByID(ID uint) error
}

type service struct {
	repo         repositoryproduct.RepositoryProduct
	repoCategory repositorycategory.RepositoryCategory
}

func New(repo repositoryproduct.RepositoryProduct, repoCategory repositorycategory.RepositoryCategory) ServiceProduct {
	return &service{repo: repo, repoCategory: repoCategory}
}

func (s *service) Create(product modelproduct.Request) (modelproduct.Response, error) {
	err := validation.ValidateProductStore(product, s.repoCategory)
	if err != nil {
		return modelproduct.Response{}, err
	}

	productEntity := new(entity.Product)
	copier.Copy(&productEntity, &product)
	createdProduct, err := s.repo.Create(*productEntity)
	if err != nil {
		return modelproduct.Response{}, err
	}

	productResponse := new(modelproduct.Response)
	copier.Copy(&productResponse, &createdProduct)
	productResponse.UpdatedAt = nil
	return *productResponse, nil
}

func (s *service) GetAll() ([]modelproduct.Response, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return []modelproduct.Response{}, err
	}

	productsResponse := make([]modelproduct.Response, len(products))
	for i, product := range products {
		productResponse := new(modelproduct.Response)
		copier.Copy(&productResponse, &product)
		productsResponse[i] = *productResponse
	}

	return productsResponse, nil
}

func (s *service) Update(product modelproduct.Request) (modelproduct.Response, error) {
	err := validation.ValidateProductStore(product, s.repoCategory)
	if err != nil {
		return modelproduct.Response{}, err
	}

	productEntity := new(entity.Product)
	copier.Copy(&productEntity, &product)
	updatedProduct, err := s.repo.Update(*productEntity)
	if err != nil {
		return modelproduct.Response{}, err
	}

	productResponse := new(modelproduct.Response)
	copier.Copy(&productResponse, &updatedProduct)
	productResponse.CreatedAt = nil
	return *productResponse, nil
}

func (s *service) DeleteByID(ID uint) error {
	err := s.repo.DeleteByID(ID)
	if err != nil {
		return err
	}
	return nil
}
