package main

import (
	"fmt"
)

type AA struct {
	xyz string
	opq string
}

func (this *AA) ChangePointer() {
	b := &AA{xyz: "xyz1", opq: "opq1"}
	this = b
	fmt.Printf("in a: %v, addr:%p, addr of pinter:%p\n", this, this, &this)
}

func (this *AA) ChangeByPointer() {
	b := AA{xyz: "xyz1", opq: "opq1"}
	*this = b
	fmt.Printf("in a: %v, addr:%p, addr of pinter:%p\n", this, this, &this)
}

func main() {
	a := &AA{xyz: "xyz0", opq: "opq0"}
	fmt.Printf("before a:%v, addr:%p, addr of pointer:%p\n", a, a, &a)

	a.ChangePointer() // <==> Change(a)
	// 调用 a.Chnage()时，Change中的this 是 指针a的一个拷贝

	fmt.Printf("after a:%v, addr:%p, addr of pointer:%p\n", a,a, &a)
}