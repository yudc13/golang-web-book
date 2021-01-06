package controller

import (
	"golangWebBook/model"
	"golangWebBook/model/response"
	"golangWebBook/service"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if 	strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		response.Failure(nil, "用户名或密码不能为空", w)
		return
	}
	user, err := service.Login(username, password)
	if err != nil {
		response.Failure(nil, err.Error(), w)
		return
	}
	response.Success(user, w)
}

func Register(w http.ResponseWriter, r *http.Request)  {
	username := r.FormValue("username")
	password := r.FormValue("password")
	rePassword := r.FormValue("rePassword")
	email := r.FormValue("email")
	if password != rePassword {
		response.Failure(nil, "两次密码不一致", w)
		return
	}
	user := &model.User{
		Username: username,
		Pwd:      password,
		Email:    email,
	}
	userId, err := service.Register(user)
	if err != nil {
		response.Failure(nil, err.Error(), w)
		return
	}
	response.Success(map[string]int64{"userId": userId}, w)
}
