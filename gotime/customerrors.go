package gotime

import "errors"

var (
	ErrInvalidLength = errors.New("time string must have at least 2 characters")
	ErrNegativeValue = errors.New("negative time values are not allowed")
	ErrInvalidUnit   = errors.New("invalid time unit")
	ErrInvalidFormat = errors.New("invalid number format")
)
