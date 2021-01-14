package controller

import (
	"golangWebBook/model"
	"golangWebBook/model/response"
	"golangWebBook/service"
	"net/http"
	"strconv"
)

func GetBookList(w http.ResponseWriter, r *http.Request) {
	//err := r.ParseForm()
	//var formData map[string]interface{}
	//json.NewDecoder(r.Body).Decode(&formData)
	//fmt.Printf("body: %+v \n", formData)
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
