package service

import "golangWebBook/model"

func GetOrders(userId int) (order model.Order, err error) {
	return model.QueryOrders(userId)
}

func AddOrder(userId int, order model.OrderInfo, books []model.Book) error {
	return model.InsertOrder(userId, order, books)
}
