/*************************************************************************
 *  @author:
 *  @date created: 2015-04-10
 *  @purpose: 1. stack 2. interface{} as any value
 ************************************************************************/
package chp01

import (
	"github.com/Oscarzhao/golang/programmingingo/chp01/stack"
	"fmt"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
