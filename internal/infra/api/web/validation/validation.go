package validation

import (
	"encoding/json"
	"errors"

	"github.com/LuisGaravaso/goexpert-auction/configs/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	valEn "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		valEn.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("invalid type error")
	} else if errors.As(validation_err, &jsonValidation) {
		causes := []rest_err.Causes{}
		for _, err := range jsonValidation {
			causes = append(causes, rest_err.Causes{
				Field:   err.Field(),
				Message: err.Translate(transl),
			})
		}
		return rest_err.NewBadRequestError("invalid field values", causes...)
	} else {
		return rest_err.NewInternalServerError("internal server error", validation_err)
	}

}
