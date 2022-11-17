package handlers

import "errors"

var (
	ErrURLExisted = errors.New("url is already used")
)
