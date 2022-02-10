package validation

import (
	"errors"

	"github.com/arfan21/golang-tokobelanja/model/modelproduct"
	"github.com/arfan21/golang-tokobelanja/repository/repositorycategory"
	validation "github.com/go-ozzo/ozzo-validation"
)

func IsCategoryExist(repo repositorycategory.RepositoryCategory) validation.RuleFunc {
	return func(value interface{}) error {
		id, ok := value.(uint)
		if !ok {
			return errors.New("invalid category id")
		}

		return repo.IsCategoryExist(id)
	}
}

func ValidateProductStore(data modelproduct.Request, repo repositorycategory.RepositoryCategory) error {
	return validation.Errors{
		"title":       validation.Validate(data.Title, validation.Required),
		"price":       validation.Validate(data.Price, validation.Required, validation.Min(0), validation.Max(50000000)),
		"stock":       validation.Validate(data.Stock, validation.Required, validation.Min(5)),
		"category_id": validation.Validate(data.CategoryID, validation.Required, validation.By(IsCategoryExist(repo))),
	}.Filter()
}
