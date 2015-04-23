package jsonutil

import (
	"errors"
)

var ErrInvalidKey = errors.New("the json key is invalid")

var ErrInvalidKeyPath = errors.New("the key path doesn't exist")

var ErrInvalidJsonObject = errors.New("the value of the key path is not a json object")

var ErrInvalidArray = errors.New("the value of the key path is not an array")

var ErrInvalidArrayIndex = errors.New("index overflow")
