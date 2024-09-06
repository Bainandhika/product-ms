package helpers

import "errors"

var ErrProductNotFound = errors.New("product not found")
var ErrProductAlreadyExists = errors.New("product already exists")
