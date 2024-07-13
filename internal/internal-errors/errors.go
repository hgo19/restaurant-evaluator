package internalerrors

import "errors"

var ErrInternal error = errors.New("Internal server error")

var NotFound error = errors.New("Not found")
