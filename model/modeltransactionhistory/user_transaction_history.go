package modeltransactionhistory

import (
	"github.com/arfan21/golang-tokobelanja/model/modelproduct"
	"github.com/arfan21/golang-tokobelanja/model/modeluser"
	"time"
)

type RequestTransaction struct {
	UserID    uint `json:"user_id,omitempty"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type ResponseMakeTransaction struct {
	Message         string `json:"message"`
	TransactionBill struct {
		TotalPrice   int    `json:"total_price"`
		Quantity     int    `json:"quantity"`
		ProductTitle string `json:"product_title"`
	} `json:"transaction_bill"`
}

type ResponseTransactionHistory struct {
	ID         uint                  `gorm:"primary_key" json:"id"`
	ProductID  uint                  `json:"product_id"`
	UserID     uint                  `json:"user_id"`
	Quantity   int                   `json:"quantity"`
	TotalPrice int                   `json:"total_price"`
	CreatedAt  time.Time             `json:"created_at"`
	Product    modelproduct.Response `json:"product"`
}

type ResponseTransactionAll struct {
	ID         uint                  `gorm:"primary_key" json:"id"`
	ProductID  uint                  `json:"product_id"`
	UserID     uint                  `json:"user_id"`
	Quantity   int                   `json:"quantity"`
	TotalPrice int                   `json:"total_price"`
	CreatedAt  time.Time             `json:"created_at"`
	Product    modelproduct.Response `json:"product"`
	User       modeluser.Response    `json:"user"`
}
