package intro

import "fmt"

func main() {
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	default:
		fmt.Println("Bigger than 9")
	}
	fmt.Print("Hello, Go")
}
