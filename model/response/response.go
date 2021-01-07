package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 1
	SUCCESS = 0
)

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
}

// 返回结果
func Result(code int, data interface{}, msg string, w http.ResponseWriter) {
	SetJsonHeader(w)
	err := json.NewEncoder(w).Encode(Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	if err != nil {
		panic(err)
	}
}

func Success(data interface{}, w http.ResponseWriter) {
	SetJsonHeader(w)
	err := json.NewEncoder(w).Encode(Response{
		Code: SUCCESS,
		Data: data,
		Msg:  "成功",
	})
	if err != nil {
		panic(err)
	}
}

func Failure(data interface{}, msg string, w http.ResponseWriter) {
	SetJsonHeader(w)
	err := json.NewEncoder(w).Encode(Response{
		Code: ERROR,
		Data: data,
		Msg:  msg,
	})
	if err != nil {
		panic(err)
	}
}
