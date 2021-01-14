package service

import "golangWebBook/model"

func GetOrders(userId int) (order model.Order, err error) {
	return model.QueryOrders(userId)
}
