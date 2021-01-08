package model

import "golangWebBook/db"

type Book struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Author string  `json:"author"`
	Sales  int     `json:"sales"`
	Stock  int     `json:"stock"`
	Cover  string  `json:"cover"`
}

type PageInfo struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"total_page"`
}

type BookPage struct {
	Books    []Book   `json:"books"`
	PageInfo PageInfo `json:"page_info"`
}

func QueryBookCountByName(bookName string) (total int64, err error) {
	query := "select count(id) as total from books where name like ?"
	row := db.Db.QueryRow(query, "%"+bookName+"%")
	err = row.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func QueryBookListByPage(bookName string, currentPage int64, pageSize int64) (books []Book, err error) {
	query := "select id, name, author, price, sales, stock, cover from books where name like ? limit ?, ?"
	offset := (currentPage - 1) * pageSize
	rows, err := db.Db.Query(query, "%"+bookName+"%", offset, pageSize)
	if err != nil {
		return []Book{}, err
	}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Cover)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books, nil
}
