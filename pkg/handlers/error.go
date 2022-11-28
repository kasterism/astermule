package handlers

import (
	"encoding/json"
)

type httpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewHttpResponse(code int, errMsg string, data interface{}) httpResponse {
	return httpResponse{
		Code: code,
		Msg:  errMsg,
		Data: data,
	}
}

func (h httpResponse) AsBytes() []byte {
	data, _ := json.Marshal(h)
	logger.Errorln("Response marshal error")
	return data
}
