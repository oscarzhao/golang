package main

import (
	"net/http"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello"
	html = html + " World"

	w.Write([]byte(html))
}
