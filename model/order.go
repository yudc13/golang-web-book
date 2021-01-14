package model

import (
	"fmt"
	"golangWebBook/db"
)

type Order struct {
	User  User   `json:"user"`
	Books []Book `json:"books"`
	Count int64  `json:"count"`
	Id    int64  `json:"id"`
}

type OrderInfo struct {
	Id     int     `json:"id"`
	Price  float64 `json:"price"`
	Count  int     `json:"count"`
	UserId int     `json:"user_id"`
}

type OrderDetail struct {
	Id      int     `json:"id"`
	BookId  int     `json:"book_id"`
	OrderId int     `json:"order_id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Author  string  `json:"author"`
	Cover   string  `json:"cover"`
}

func QueryOrders(userId int) (o Order, err error) {
	query := `SELECT
		users.id AS user_id,
		users.username,
		users.email,
		orders.id AS order_id,
		orders.count AS order_count,
        order_details.id as book_id,
		order_details.name,
		order_details.price,
		order_details.author,
		order_details.cover 
	FROM
		orders
		JOIN users ON users.id = orders.user_id
		JOIN order_details ON orders.id = order_details.order_id
	WHERE users.id = ?`
	rows, err := db.Db.Query(query, userId)
	if err != nil {
		return Order{}, err
	}
	order := Order{
		User: User{
			Id:       0,
			Username: "",
			Email:    "",
		},
		Books: nil,
		Count: 0,
		Id:    0,
	}
	var books []Book
	for rows.Next() {
		book := Book{
			Id:     0,
			Name:   "",
			Price:  0,
			Author: "",
			Cover:  "",
		}
		err := rows.Scan(&order.User.Id, &order.User.Username, &order.User.Email, &order.Id, &order.Count, &book.Id, &book.Name, &book.Price, &book.Author, &book.Cover)
		if err != nil {
			return Order{}, err
		}
		books = append(books, book)
	}
	order.Books = books
	return order, nil
}

func InsertOrder(userId int, order OrderInfo, books []Book) error {
	// 开启事务
	tx, err := db.Db.Begin()
	if err != nil {
		return err
	}
	orderSql := "INSERT INTO orders (price, count, user_id) VALUES (?, ?, ?)"
	result, err := tx.Exec(orderSql, order.Price, order.Count, userId)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	orderDetailSql := "INSERT INTO order_details (order_id, book_id, price, name, author, cover) VALUES "
	var values []interface{}
	for _, b := range books {
		orderDetailSql += "(?, ?, ?, ?, ?, ?),"
		values = append(values, lastInsertId, b.Id, b.Price, b.Name, b.Author, b.Cover)
	}
	// 去掉最后一个,
	orderDetailSql = orderDetailSql[0 : len(orderDetailSql)-1]
	_, e := tx.Exec(orderDetailSql, values...)
	if e != nil {
		err := tx.Rollback()
		fmt.Println("add orderDetail err: ", e, err)
		return e
	}
	err = tx.Commit()
	if err != nil {
		err := tx.Rollback()
		return err
	}
	return nil
}
