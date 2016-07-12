package landmines

import (
	"fmt"
	"testing"
)

/*
It's because t.PrintID() is equivalent to

f := (*T).PrintID
f(t)
or, more compactly

(*T).PrintID(t)
*/
type T struct {
	ID int
}

func (t *T) PrintID() {
	fmt.Println(t.ID)
}

func F1() {
	ts := []T{{1}, {2}, {3}}
	for _, t := range ts {
		defer t.PrintID()
	}
}

func F2() {
	ts := []*T{&T{1}, &T{2}, &T{3}}
	for _, t := range ts {
		defer t.PrintID()
	}
}

func TestPassByValues(t *testing.T) {
	fmt.Println("F1()")
	F1()
	fmt.Println()
	fmt.Println("F2()")
	F2()
}
