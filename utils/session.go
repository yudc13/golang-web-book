package utils

import "github.com/gorilla/sessions"

var SessionStore = sessions.NewCookieStore([]byte("web_book_server"))
