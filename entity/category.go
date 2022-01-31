package entity

import "time"

type Category struct {
	ID                uint      `json:"id gorm:primaryKey"`
	Type              string    `json:"type" gorm:"not null"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Products          []Product `json:"products"`
}
