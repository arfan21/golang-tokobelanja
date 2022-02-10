package modelcategory

import (
	"time"

	"github.com/arfan21/golang-tokobelanja/model/modelproduct"
)

type Request struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
}

type Response struct {
	ID                uint                    `json:"id"`
	Type              string                  `json:"type"`
	SoldProductAmount int                     `json:"sold_product_amount"`
	CreatedAt         *time.Time              `json:"created_at,omitempty"`
	UpdatedAt         *time.Time              `json:"updated_at,omitempty"`
	Products          []modelproduct.Response `json:"products"`
}
