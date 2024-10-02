package schemavalidator

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type SchemaValidator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
}

func New() *SchemaValidator {
	en := en.New()
	uni := ut.New(en, en)

	validate := validator.New(validator.WithRequiredStructEnabled())

	// Register translations for translators
	schemaValidator := SchemaValidator{validate, uni}

	schemaValidator.registerDefaultTranslations()

	return &schemaValidator
}

func (schemaValidator *SchemaValidator) BindJSON(req *http.Request, result any) error {
	err := json.NewDecoder(req.Body).Decode(result)

	if err != nil {
		return customerrors.ErrJSONDecoding
	}

	value := reflect.ValueOf(result)

	switch value.Kind() {
	case reflect.Ptr:
		err := schemaValidator.validate.Struct(value.Elem().Interface())
		return err
	case reflect.Struct:
		err := schemaValidator.validate.Struct(result)
		return err
	}

	return errors.New("invalid result type")
}

func (schemaValidator *SchemaValidator) GetTranslationErrors(errors validator.ValidationErrors, locale string) []models.ErrorObject {
	trans, _ := schemaValidator.uni.GetTranslator("en")

	translationErrors := errors.Translate(trans)

	errorObjects := make([]models.ErrorObject, len(translationErrors))

	i := 0
	for _, message := range translationErrors {
		errorObjects[i] = models.ValidationError(message)
		i++
	}

	return errorObjects
}

func (schemaValidator *SchemaValidator) registerDefaultTranslations() {
	// TODO: Handle more locales
	enTranslator, _ := schemaValidator.uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(schemaValidator.validate, enTranslator)
}
