package validation

import (
	"github.com/arfan21/golang-tokobelanja/model/modelcategory"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateCategoryStore(data modelcategory.Request) error {
	return validation.Errors{
		"type": validation.Validate(data.Type, validation.Required),
	}.Filter()
}
