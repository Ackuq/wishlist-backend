package schemavalidator

import (
	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
)

func Init() {
	en := en.New()

	uni = ut.New(en, en)
	validate = validator.New(validator.WithRequiredStructEnabled())

	registerDefaultTranslations()
}

func ValidateStruct(obj any) error {
	return validate.Struct(obj)
}

func GetTranslationErrors(errors validator.ValidationErrors, locale string) []models.ErrorObject {
	trans, _ := uni.GetTranslator("en")

	translationErrors := errors.Translate(trans)

	errorObjects := make([]models.ErrorObject, len(translationErrors))

	i := 0
	for _, message := range translationErrors {
		errorObjects[i] = models.ValidationError(message)
		i++
	}

	return errorObjects
}

func registerDefaultTranslations() {
	// TODO: Handle more locales
	enTranslator, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, enTranslator)
}
