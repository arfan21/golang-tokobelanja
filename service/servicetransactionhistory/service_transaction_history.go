package servicetransactionhistory

import (
	"github.com/arfan21/golang-tokobelanja/constant"
	"github.com/arfan21/golang-tokobelanja/entity"
	"github.com/arfan21/golang-tokobelanja/model/modeltransactionhistory"
	"github.com/arfan21/golang-tokobelanja/repository/repositorytransactionhistory"
	"github.com/arfan21/golang-tokobelanja/validation"
	"github.com/jinzhu/copier"
)

type ServiceTransactionHistory interface {
	CreateTransactionHistory(request modeltransactionhistory.RequestTransaction) (modeltransactionhistory.ResponseMakeTransaction, error)
	GetTransactionHistory(userID uint) ([]modeltransactionhistory.ResponseTransactionHistory, error)
	GetTransactionHistories() ([]modeltransactionhistory.ResponseTransactionAll, error)
}

type Service struct {
	repo repositorytransactionhistory.RepositoryTransactionHistory
}

func (s *Service) GetTransactionHistory(userID uint) ([]modeltransactionhistory.ResponseTransactionHistory, error) {
	transactionHistories, err := s.repo.GetTransactionHistory(userID)
	if err != nil {
		return []modeltransactionhistory.ResponseTransactionHistory{}, err
	}
	var responseTransactionHistories []modeltransactionhistory.ResponseTransactionHistory
	copier.Copy(&responseTransactionHistories, &transactionHistories)
	return responseTransactionHistories, nil
}

func (s *Service) GetTransactionHistories() ([]modeltransactionhistory.ResponseTransactionAll, error) {
	transactionHistories, err := s.repo.GetAllTransactionHistories()
	if err != nil {
		return []modeltransactionhistory.ResponseTransactionAll{}, err
	}
	var responseTransactionHistories []modeltransactionhistory.ResponseTransactionAll
	copier.Copy(&responseTransactionHistories, &transactionHistories)
	return responseTransactionHistories, nil
}

func (s *Service) CreateTransactionHistory(request modeltransactionhistory.RequestTransaction) (modeltransactionhistory.ResponseMakeTransaction, error) {

	resp := modeltransactionhistory.ResponseMakeTransaction{}

	err := validation.ValidateMakeTransactions(request)
	if err != nil {
		return resp, err
	}

	// stock availability
	product, err := s.repo.GetProduct(request.ProductID)
	if err != nil {
		return resp, err
	}

	// validate stock
	if product.Stock < request.Quantity {
		return resp, constant.ErrorOutOfStock
	}

	// validate balance
	user, err := s.repo.GetUserByID(request.UserID)
	if err != nil {
		return resp, err
	}

	totalPrice := product.Price * request.Quantity

	if user.Balance < totalPrice {
		return resp, constant.ErrorBalance
	}

	transactionHistory := entity.TransactionHistory{}
	copier.Copy(&transactionHistory, &request)
	transactionHistory.TotalPrice = totalPrice
	_, err = s.repo.CreateTransaction(transactionHistory)

	if err != nil {
		return resp, err
	}

	resp.Message = "you successfully purchase the product"
	resp.TransactionBill.TotalPrice = totalPrice
	resp.TransactionBill.Quantity = request.Quantity
	resp.TransactionBill.ProductTitle = product.Title
	return resp, nil
}

func New(repo repositorytransactionhistory.RepositoryTransactionHistory) ServiceTransactionHistory {
	return &Service{repo: repo}
}
