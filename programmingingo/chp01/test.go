package chp01

import (
	"fmt"
	"regexp"
)

func main() {
	ss := "1234"
	for _, c := range ss {
		fmt.Println(c)
	}
	ssList := []string{"aa:bb", "cc:dd", "ee:ff"}
	for index, content := range ssList {
		fmt.Printf("index: %d, content: %s\n", index, content)
	}

	// Compile the expression once, usually at init time
	// Use raw strings to avoid having to quote the backslashes
	var validId = regexp.MustCompile(`^([a-z]+\[[0-9]+\])$`)
	fmt.Println(validId.MatchString("adam[23]"))  // true
	fmt.Println(validId.MatchString("eve[7]"))    // true
	fmt.Println(validId.MatchString("oscar[0]d")) // false
	fmt.Println(validId.MatchString("Job[48]"))   // false
	fmt.Println(validId.MatchString("snakey"))    // false

	var wordPatt = regexp.MustCompile(`[., ]*([A-Za-z]+)[., ]*`)
	dst := `Hello!
My name is Oscar.  My hometown is Henan.
There are five members in my family: my father, mother, younger sister, younger brother and me.`
	fmt.Println(dst)
	for i, c := range wordPatt.FindAllStringSubmatch(dst, -1) {
		fmt.Printf("i=%d, c=%q\n", i, c)
	}

}
