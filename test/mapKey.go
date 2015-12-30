package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	id2Person := make(map[string]Person, 5)

	id2Person["12324"] = Person{Name: "Oscar", Age: "24"}

	fmt.Printf("map: %v\n", id2Person)

	key := "4321"
	person, found := id2Person[key]
	if found {
		fmt.Printf("key \"%s\" found, value:%v\n", key, person)
	} else {
		fmt.Printf("key \"%s\" not found\n", key)
	}

}
