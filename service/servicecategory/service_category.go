package servicecategory

import (
	"github.com/arfan21/golang-tokobelanja/entity"
	"github.com/arfan21/golang-tokobelanja/model/modelcategory"
	"github.com/arfan21/golang-tokobelanja/repository/repositorycategory"
	"github.com/arfan21/golang-tokobelanja/validation"
	"github.com/jinzhu/copier"
)

type ServiceCategory interface {
	Create(request modelcategory.Request) (modelcategory.Response, error)
	GetAll() ([]modelcategory.Response, error)
	Update(request modelcategory.Request) (modelcategory.Response, error)
	Delete(id uint64) error
}

type service struct {
	repo repositorycategory.RepositoryCategory
}

func New(repo repositorycategory.RepositoryCategory) ServiceCategory {
	return &service{repo: repo}
}

func (s *service) Delete(id uint64) error {
	err := s.repo.Delete(uint(id))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(request modelcategory.Request) (modelcategory.Response, error) {
	err := validation.ValidateCategoryStore(request)
	if err != nil {
		return modelcategory.Response{}, err
	}

	entityCategory := new(entity.Category)
	copier.Copy(entityCategory, &request)
	update, err := s.repo.Update(*entityCategory)
	if err != nil {
		return modelcategory.Response{}, err
	}

	resp := new(modelcategory.Response)
	copier.Copy(resp, &update)
	resp.CreatedAt = nil
	return *resp, nil
}

func (s *service) GetAll() ([]modelcategory.Response, error) {
	gets, err := s.repo.GetAll()
	if err != nil {
		return []modelcategory.Response{}, err
	}

	var resp []modelcategory.Response
	copier.Copy(&resp, &gets)
	return resp, nil
}

func (s *service) Create(request modelcategory.Request) (modelcategory.Response, error) {
	err := validation.ValidateCategoryStore(request)
	if err != nil {
		return modelcategory.Response{}, err
	}

	entityCategory := new(entity.Category)
	copier.Copy(entityCategory, &request)
	create, err := s.repo.Create(*entityCategory)
	if err != nil {
		return modelcategory.Response{}, err
	}

	resp := new(modelcategory.Response)
	copier.Copy(resp, &create)
	resp.UpdatedAt = nil
	return *resp, nil
}
