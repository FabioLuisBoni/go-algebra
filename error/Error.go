package algebra_error

import "errors"

var (
	ErrMissingPoints     error = errors.New("missing points")
	ErrMissingExpression error = errors.New("missing expression")
)
