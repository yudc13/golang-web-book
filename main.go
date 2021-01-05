package main

import (
	"fmt"
	"golangWebBook/db"
	"html"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(w, "数据库链接成功 %q", html.EscapeString(r.URL.Path))
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
