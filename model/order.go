package model

import "golangWebBook/db"

type Order struct {
	User  User   `json:"user"`
	Books []Book `json:"books"`
	Count int64  `json:"count"`
	Id    int64  `json:"id"`
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
