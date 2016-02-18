package main

/*
	An example to create Race Conditions in map[]
	from: https://github.com/google/sanitizers/wiki/ThreadSanitizerGoManual
	detect race conditions: go run -race xxx.go
*/

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
