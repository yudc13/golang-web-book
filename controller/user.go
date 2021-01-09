package controller

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"golangWebBook/model"
	"golangWebBook/model/response"
	"golangWebBook/service"
	"net/http"
	"os"
	"strings"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("user")))

func GetUserInfo(w http.ResponseWriter, r *http.Request)  {
	session, err := store.Get(r, "user")
	if err != nil {
		response.Failure(nil, err.Error(), w)
		return
	}
	user := session.Values["user"]
	response.Success(user, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		response.Failure(nil, "用户名或密码不能为空", w)
		return
	}
	user, err := service.Login(username, password)
	if err != nil {
		response.Failure(nil, err.Error(), w)
		return
	}
	// 保存用户信息到session中
	session, _ := store.Get(r, "user")
	gob.Register(user)
	session.Values["user"] = user
	err = session.Save(r, w)
	if err !=  nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Success(user, w)
}

func Register(w http.ResponseWriter, r *http.Request) {
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
