package modelproduct

import "time"

type Request struct {
	ID         uint   `json:"id,omitempty"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_id"`
}

type Response struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	Price      int        `json:"price"`
	Stock      int        `json:"stock"`
	CategoryID uint       `json:"category_id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
