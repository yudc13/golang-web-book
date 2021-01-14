package controller

import (
	"golangWebBook/model"
	"golangWebBook/model/response"
	"golangWebBook/service"
	"golangWebBook/utils"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	session, err := utils.SessionStore.Get(r, "user")
	if err != nil {
		response.Failure(nil, "获取用户信息失败", w)
		return
	}
	user := session.Values["user"]
	u := user.(model.User)
	order, err := service.GetOrders(u.Id)
	response.Success(order, w)
}
