package main

/*
	The code is used to demonstrate the mechanism and techniques of reflection
	derived from: http://blog.golang.org/laws-of-reflection
*/

import (
	"fmt"
	"reflect"
)

/*
	the Kind of a reflection object describes the underlying type, not the static type.
	If a reflection object contains a value of a user-defined integer type, as in the following example
	the Kind of v is still reflect.Int, even though the static type of x is MyInt, not int.
	In other words, the Kind cannot discriminate an int from a MyInt even though the Type can.
*/
func reflectLaws3() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	// OUTPUT
	// v.Kind(): int, v.Type(): main.MyInt
	fmt.Printf("v.Kind(): %s, v.Type(): %s\n", v.Kind(), v.Type())
}

func reflectStruct() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func main() {
	reflectStruct()
}
