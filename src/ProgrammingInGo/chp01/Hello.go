/*************************************************************************
 * @author:
 * @date created: 2015-04-10
 * @purpose:      1. read cmd params; 2. strings slice; 3. fmt.Println
 ************************************************************************/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
