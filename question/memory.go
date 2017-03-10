// m and n are of different capacities
// and Why?
// once you print it with `%v`, the two byte array are of the same length
// and Why?
package main

import (
	"fmt"
)

func main() {
	m := []byte("acaaaaaaaaaaaaaa")
	n := []byte("acaaaaaaaaaaaaaa")

	fmt.Printf("len(n): %d, cap(n):%d\n", len(n), cap(n))
	fmt.Printf("len(m): %d, cap(m):%d\n", len(m), cap(m))

	fmt.Printf("copy...\n")
	copy(n, m)
	fmt.Printf("len(n): %d, cap(n):%d\n", len(n), cap(n))
	fmt.Printf("len(m): %d, cap(m):%d\n", len(m), cap(m))

	fmt.Printf("n: %s\n", n)
}
