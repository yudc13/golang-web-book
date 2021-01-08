package service

import (
	"fmt"
	"testing"
)

func TestGetBookPage(t *testing.T)  {
	bookPage, _ := GetBookList("", 1, 3)
	fmt.Printf("books: %+v \n", bookPage.Books)
	fmt.Printf("pageInfo: %+v \n", bookPage.PageInfo)
}
