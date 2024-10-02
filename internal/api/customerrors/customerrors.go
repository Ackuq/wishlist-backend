package customerrors

import "errors"

var ErrJSONDecoding = errors.New("could not decode json into desired type")
