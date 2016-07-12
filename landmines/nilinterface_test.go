package landmines

import (
	"fmt"
	"testing"
)

// Cat Nil interface is not the same as having a nil pointer in the interface.
type Cat interface {
	Meow()
}

type Tabby struct{}

func (*Tabby) Meow() { fmt.Println("meow") }

func GetACat() Cat {
	var myTabby *Tabby
	// Oops, we forgot to set myTabby to a real value
	return myTabby
}

func GetArray() (arr []string) {
	return
}

// http://golang.org/doc/faq#nil_error
// https://gist.github.com/lavalamp/4bd23295a9f32706a48f
func TestGetACat(t *testing.T) {
	if GetACat() == nil {
		t.Errorf("Forgot to return a real cat!")
	} else {
		t.Logf("should be not nil")
	}
}

func TestGetArray(t *testing.T) {
	if GetArray() == nil {
		t.Errorf("Forgot to return a real array!")
	} else {
		t.Logf("should be not nil")
	}
}
