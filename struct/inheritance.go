package main

/**
*    Title: 继承非暴露属性
*    Source: https://talks.golang.org/2015/go4cpp.slide#39
*    Aim: 隐藏一些字段，只暴露方法
*
 */

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
	myown int
}

type Employee struct {
	Person
}

func main() {
	e := Employee{Person{"John", "Doe", 49, 1}}
	fmt.Printf("e:%v\n", e.myown) // 暴露属性可以继承
}
