package intro

import "fmt"

// this is a comment

func main() {
	a, b := fmt.Println("My name is Oscar")
	// equals to
	// var a, b = fmt.Println("Let's go")
	fmt.Println(a, b)

	var X string
	X = "Hello World"
	fmt.Println(X)

	fmt.Println("1 == 1:", 1 == 1)

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
