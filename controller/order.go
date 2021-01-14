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

func AddOrder(w http.ResponseWriter, r *http.Request)  {
	session, err := utils.SessionStore.Get(r, "user")
	if err != nil {
		response.Failure(nil, "获取用户信息失败", w)
		return
	}
	user := session.Values["user"]
	u := user.(model.User)
	order := model.OrderInfo{
		Price:  300,
		Count:  3,
		UserId: u.Id,
	}
	books := []model.Book{
		{
			Id:     10,
			Name:   "Go程序设计语言",
			Price:  126.00,
			Author: "艾伦 A.A.多诺万",
			Cover:  "http://img3m4.ddimg.cn/67/29/1824606004-1_w_1.jpg",
		},
		{
			Id:     200,
			Name:   "龙猫",
			Price:  49.00,
			Author: "宫崎骏",
			Cover:  "http://img3m7.ddimg.cn/74/31/29174087-1_w_4.jpg",
		},
	}
	err = service.AddOrder(u.Id, order, books)
	if err != nil {
		response.Failure(nil, "添加订单失败: " + err.Error(), w)
		return
	}
	response.Success(nil, w)
}
