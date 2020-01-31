package models

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models/consts"
)

type DataReturns struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *DataReturns) String() string{
	dat, _ := json.Marshal(r)
	return string(dat)
}

func WithSuccess(msg string, data interface{}) *DataReturns{
	return &DataReturns{
		Code:    consts.CodeOK,
		Message: msg,
		Data:    data,
	}
}

func WithError(code int, msg string) *DataReturns{
	return &DataReturns{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}
