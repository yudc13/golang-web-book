package controller

import (
	"golangWebBook/model"
	"golangWebBook/model/response"
	"golangWebBook/service"
	"net/http"
	"strconv"
)

func GetBookList(w http.ResponseWriter, r *http.Request) {
	bookName := r.PostFormValue("bookName")
	currentPage := r.PostFormValue("currentPage")
	pageSize := r.PostFormValue("pageSize")
	page, _ := strconv.ParseInt(currentPage, 10, 64)
	size, _ := strconv.ParseInt(pageSize, 10, 64)
	bookPage, err := service.GetBookList(bookName, page, size)
	if err != nil {
		response.Failure([]model.Book{}, err.Error(), w)
	}
	response.Success(bookPage, w)

}
