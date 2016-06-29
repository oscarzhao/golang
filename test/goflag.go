package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Printf("commandline:%v\n\n\n", flag.CommandLine)
	fmt.Printf("args:%v\n", flag.Args())

	flag.Parse()
	fmt.Printf("commandline:%v\n", flag.CommandLine)
	fmt.Printf("args:%v\n", flag.Args())
}
