package common

import (
	"fmt"
)

type ParseError struct {
	HTTPStatusCode int
	Err            error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Parse data error, status code: %d, message: %s", e.HTTPStatusCode, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}
