package model 

import (
	"errors"
)

var (
	ErrInvalidRomanNumber      = errors.New("Request number is in invalid format")
	ErrParamNotFound      = errors.New("I have no idea what you are talking about")
	ErrItemNotFound		= errors.New("I have no idea what you are talking about")
	ErrInvalidCommand		= errors.New("I have no idea what you are talking about")
)