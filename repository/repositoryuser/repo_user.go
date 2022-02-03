package repositoryuser

import (
	"errors"

	"github.com/arfan21/golang-tokobelanja/constant"
	"github.com/arfan21/golang-tokobelanja/entity"
	"gorm.io/gorm"
)

type RepositoryUser interface {
	Create(data entity.User) (entity.User, error)
	IsEmailExist(email string) error
	Login(email string) (entity.User, error)
	Update(data entity.User) (entity.User, error)
	GetByID(id uint) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryUser {
	return &repository{db: db}
}

func (r *repository) Create(data entity.User) (entity.User, error) {
	if data.Role != constant.AdminRole && data.Role != constant.MemberRole {
		return entity.User{}, constant.ErrorInvalidRole
	}
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.User{}, err
	}

	return data, nil
}

func (r *repository) IsEmailExist(email string) error {
	user := new(entity.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	return constant.ErrorEmailAlreadyExists
}

func (r *repository) Login(email string) (entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}

func (r *repository) Update(data entity.User) (entity.User, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return entity.User{}, err
	}

	return data, nil
}

func (r *repository) GetByID(id uint) (entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}
