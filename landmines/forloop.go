package landmines

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Println("1111", i)
		defer func() { fmt.Println("2222", i) }()
		defer func(i int) { fmt.Println("33333", i) }(i)
		defer print(&i)
		go fmt.Println("44444", i)
		go func() { fmt.Println("55555", i) }()
	}
}

func print(pi *int) { fmt.Println("ppppp", *pi) }
