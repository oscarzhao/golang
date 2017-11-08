package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var dir string

func init() {
	flag.StringVar(&dir, "dir", "./Downloads/", "file directory to expose")
	flag.Parse()
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("get interface addresses fails, err=%s\n", err)
	}
	for _, addr := range addrs {
		if strings.HasPrefix(addr.String(), "192.168.1.") {
			fmt.Printf("%s, %s\n", addr.Network(), addr.String())
			return strings.TrimSuffix(addr.String(), "/24")
		}
	}
	return "0.0.0.0"
}

func main() {
	ip := getLocalIP()
	// To serve a directory on disk (/tmp) under an alternate URL
	// path (/tmpfiles/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:it
	handler := http.StripPrefix("/", http.FileServer(http.Dir(dir)))
	addr := fmt.Sprintf("%s:8080", ip)
	log.Printf("Start listening on %s\n File Directory: %s\n", addr, dir)
	log.Fatal(http.ListenAndServe(addr, handler))
}
