package main

import (
	"net/http"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello"
	html = html + " World"

	w.Write([]byte(html))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHelloWorld)

	http.ListenAndServe(":8080", mux)
}
