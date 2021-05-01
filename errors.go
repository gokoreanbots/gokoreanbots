package gokoreanbots

import "errors"

var (
	ErrTooManyRequests = errors.New("429 too many request")
	ErrUnauthorized    = errors.New("410 unauthorized")
	ErrBadRequest      = errors.New("400 bad request")
)
