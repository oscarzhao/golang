package main

/**
*    Title: interface and composition
*    Source: https://talks.golang.org/2015/go4cpp.slide#39
*    Aim: 隐藏一些字段，只暴露方法
*
*   Struct embedding of interfaces
    Embedding an interface:
    1. more types can be used
    2. limits what is added to the embedding type
*/

import (
	"fmt"
)

type Namer interface {
	Name() string
}

type Person struct {
	First string
	Last  string
	Age   int
}

func (e Person) Name() string { return e.First + e.Last }

type Employee struct {
	Namer
}

func main() {
	e := Employee{Person{"John", "Doe", 49}}
	fmt.Printf("e:%v\n", e)
  
}
