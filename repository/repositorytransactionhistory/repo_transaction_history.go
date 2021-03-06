package repositorytransactionhistory

import (
	"github.com/arfan21/golang-tokobelanja/entity"
	"gorm.io/gorm"
)

type RepositoryTransactionHistory interface {
	CreateTransaction(data entity.TransactionHistory) (entity.TransactionHistory, error)
	GetUserByID(userID uint) (entity.User, error)
	GetProduct(productID uint) (entity.Product, error)

	GetTransactionHistory(userID uint) ([]entity.TransactionHistory, error)
	GetAllTransactionHistories() ([]entity.TransactionHistory, error)
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) GetTransactionHistory(userID uint) ([]entity.TransactionHistory, error) {
	var transactionHistories []entity.TransactionHistory
	err := r.db.Where("user_id = ?", userID).Preload("Product").
		Find(&transactionHistories).Error
	if err != nil {
		return transactionHistories, err
	}
	return transactionHistories, nil
}

func (r *Repository) GetAllTransactionHistories() ([]entity.TransactionHistory, error) {
	var transactionHistories []entity.TransactionHistory
	err := r.db.Preload("Product").Preload("User").
		Find(&transactionHistories).Error
	if err != nil {
		return transactionHistories, err
	}
	return transactionHistories, nil
}

func (r *Repository) GetUserByID(userID uint) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repository) GetProduct(productID uint) (entity.Product, error) {
	var product entity.Product
	err := r.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *Repository) getCategory(id uint) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *Repository) CreateTransaction(data entity.TransactionHistory) (entity.TransactionHistory, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	// update product stock
	product, err := r.GetProduct(data.ProductID)
	if err != nil {
		return entity.TransactionHistory{}, err
	}
	product.Stock = product.Stock - data.Quantity
	err = r.db.Updates(&product).Error
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	// update balance
	user, err := r.GetUserByID(data.UserID)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	user.Balance = user.Balance - data.TotalPrice
	err = r.db.Model(&user).Where("id = ?", user.ID).Update("balance", user.Balance).Error
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	// update category sold
	category, err := r.getCategory(product.CategoryID)
	if err != nil {
		return entity.TransactionHistory{}, err
	}
	category.SoldProductAmount = category.SoldProductAmount + data.Quantity
	err = r.db.Updates(&category).Error
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	return data, nil
}

func New(db *gorm.DB) RepositoryTransactionHistory {
	return &Repository{db}
}
