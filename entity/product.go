package entity

import "time"

type Product struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
