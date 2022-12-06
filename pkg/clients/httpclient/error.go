package httpclient

import "errors"

var (
	ErrRequest = errors.New("request error")
	ErrAction  = errors.New("undifined action")
)
