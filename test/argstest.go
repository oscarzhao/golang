package main

import (
	"fmt"
)

func func1(action string, words ...string) {
	fmt.Printf("func1 pass params \"%#v\" to func2\n....\n", words)
	wordInterface := make([]interface{}, 0, 8)
	for _, word := range words {
		wordInterface = append(wordInterface, word)
	}
	func2("from func1", wordInterface...)
}

func func2(src string, args ...interface{}) {
	fmt.Printf("got action from \"%s\", content: %#v\n", src, args)
}

func main() {
	func1("say", "goodbye", "my", "fair", "lady")
}
