package parser

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Parser struct {
	validate *validator.Validate
}

func New() *Parser {
	validate := validator.New()

	return &Parser{validate: validate}
}

func (parser Parser) BindJSON(req *http.Request, result any) error {
	err := json.NewDecoder(req.Body).Decode(result)

	if err != nil {
		// TODO: Better handle decoding errors
		return err
	}

	value := reflect.ValueOf(result)

	switch value.Kind() {
	case reflect.Ptr:
		err := parser.validate.Struct(value.Elem().Interface())
		return err
	case reflect.Struct:
		err := parser.validate.Struct(result)
		return err
	}

	return errors.New("invalid result type")
}
