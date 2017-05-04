package main

import (
	"fmt"
	"runtime"
)

func main() {
	memoryLeaking()
}
func memoryLeaking() {
	type T struct {
		v [1 << 20]int
		t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T
	// SetFinalizer will make x escape to heap.
	// Following two lines combined will make
	// x and y not collectable.
	x.t, y.t = &y, &x // y also escapecapes to heap.
	runtime.SetFinalizer(&x, finalizer)
}
