package model

import (
	"fmt"
	"testing"
)

func TestQueryBookByPage(t *testing.T) {
	books, _ := QueryBookListByPage("", 1, 6)
	fmt.Println(books)
}

func TestQueryBookCountByName(t *testing.T) {
	total, _ := QueryBookCountByName("")
	fmt.Println(total)
}
