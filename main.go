package main

import (
	"golangWebBook/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/books", controller.GetBookList)
	http.HandleFunc("/user", controller.GetUserInfo)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
