package service

import (
	"golangWebBook/model"
)

func GetBookList(bookName string, currentPage int64, pageSize int64) (bookPage model.BookPage, err error) {
	books, err := model.QueryBookListByPage(bookName, currentPage, pageSize)
	total, _ := model.QueryBookCountByName(bookName)
	if err != nil {
		return model.BookPage{}, err
	}
	var totalPage int64
	if total % pageSize == 0 {
		totalPage = total / pageSize
	} else {
		totalPage = total / pageSize + 1
	}
	bookPage = model.BookPage{
		Books: books,
		PageInfo: model.PageInfo{
			CurrentPage: currentPage,
			PageSize:    pageSize,
			Total:       total,
			TotalPage:   totalPage,
		},
	}
	return bookPage, nil
}
