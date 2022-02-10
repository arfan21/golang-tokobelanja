package validation

import (
	"github.com/arfan21/golang-tokobelanja/model/modeltransactionhistory"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateMakeTransactions(data modeltransactionhistory.RequestTransaction) error {
	return validation.Errors{
		"product_id": validation.Validate(data.ProductID, validation.Required),
		"quantity":   validation.Validate(data.Quantity, validation.Required, validation.Min(1)),
	}.Filter()
}
