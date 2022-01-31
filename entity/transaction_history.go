package entity

import "time"

type TransactionHistory struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	ProductID  uint      `json:"product_id"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	UserID     uint      `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
