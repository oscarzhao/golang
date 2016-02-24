package main

/*
  test case for switch structure (NOT equal to that of C++)
  ```
  switch val {
    case "111":
    case "222":
      do something
  }
  ```
*/

import (
	"fmt"
)

func main() {
	var testcases = []int{1, 2, 3}
	for _, a := range testcases {
		switch a {
		case 1:
		case 2:
			fmt.Printf("a=%d (value is 1 or 2)\n", a)
		default:
			fmt.Printf("a=%d (Other value)\n", a)
		}
	}
}
