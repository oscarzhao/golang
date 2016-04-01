package main

import (
	"fmt"
	"runtime"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer recoverContent()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func recoverContent() {
	fmt.Println("calling recoverContent")
	if r := recover(); r != nil {
		fmt.Printf("Internal error: %v", r)
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		fmt.Printf("%s\n", string(buf[0:stackSize]))
	}
}
