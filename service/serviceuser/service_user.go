package serviceuser

import (
	"github.com/arfan21/golang-tokobelanja/constant"
	"github.com/arfan21/golang-tokobelanja/entity"
	"github.com/arfan21/golang-tokobelanja/helper"
	"github.com/arfan21/golang-tokobelanja/model/modeluser"
	"github.com/arfan21/golang-tokobelanja/repository/repositoryuser"
	"github.com/arfan21/golang-tokobelanja/validation"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	Create(data modeluser.Request) (modeluser.Response, error)
	Login(data modeluser.RequestLogin) (modeluser.ResponseLogin, error)
	Update(data modeluser.RequestTopUp) (modeluser.Response, error)
}

type service struct {
	repo repositoryuser.RepositoryUser
}

func New(repo repositoryuser.RepositoryUser) ServiceUser {
	return &service{repo: repo}
}

func (s *service) Create(data modeluser.Request) (modeluser.Response, error) {
	err := validation.ValidateUserCreate(data, s.repo)
	if err != nil {
		return modeluser.Response{}, err
	}

	entityUser := new(entity.User)

	copier.Copy(&entityUser, &data)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(entityUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return modeluser.Response{}, err
	}
	entityUser.Password = string(hashedPassword)
	entityUser.Role = constant.MemberRole

	createdUser, err := s.repo.Create(*entityUser)
	if err != nil {
		return modeluser.Response{}, err
	}

	resp := modeluser.Response{}

	copier.Copy(&resp, &createdUser)
	resp.UpdatedAt = nil

	return resp, nil
}

func (s *service) Login(data modeluser.RequestLogin) (modeluser.ResponseLogin, error) {
	err := validation.ValidateUserLogin(data)
	if err != nil {
		return modeluser.ResponseLogin{}, err
	}

	dataUser, err := s.repo.Login(data.Email)
	if err != nil {
		return modeluser.ResponseLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(data.Password))
	if err != nil {
		return modeluser.ResponseLogin{}, constant.ErrorInvalidLogin
	}

	token, err := helper.NewJwt(dataUser.ID, dataUser.Role)
	if err != nil {
		return modeluser.ResponseLogin{}, err
	}

	resp := modeluser.ResponseLogin{}
	resp.Token = token

	return resp, nil
}

func (s *service) Update(data modeluser.RequestTopUp) (modeluser.Response, error) {
	err := validation.ValidateUserUpdate(data)
	if err != nil {
		return modeluser.Response{}, err
	}

	dataUser, err := s.repo.GetByID(data.ID)
	if err != nil {
		return modeluser.Response{}, err
	}

	dataUser.Balance = dataUser.Balance + data.Balance

	updatedUser, err := s.repo.Update(dataUser)
	if err != nil {
		return modeluser.Response{}, err
	}

	resp := modeluser.Response{}

	copier.Copy(&resp, &updatedUser)
	resp.CreatedAt = nil

	return resp, nil
}
