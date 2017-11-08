package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "server host")
	flag.IntVar(&port, "port", 8080, "server port")
}

func parse(w http.ResponseWriter, r *http.Request) {

	resp := struct {
		UserAgent  string `json:"userAgent"`
		RequestURI string `json:"requestURI"`
		RemoteAddr string `json:"remoteAddr"`
	}{
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		RemoteAddr: r.RemoteAddr,
	}

	raw, _ := json.Marshal(resp)

	w.Write(raw)
	// w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", parse)
	serverPort := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatalf("listen %s fails, err=%s\n", serverPort, err)
	}
	log.Printf("Start listening %s\n", serverPort)
}
