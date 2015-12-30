/*************************************************************************
 *  @author:  http://www.zhihu.com/question/27161493
 *  @date created:  2015-04-10
 *  @purpose:  1. strange behavior of slice, 2. diff between slice,array
 *  @explanation:  blog.golang.org/go-slices-usage-and-internals
 ************************************************************************/
package intro

import (
	"fmt"
)

func main() {
	arr := [5]int{}                 // arr is an array with cap of 5, default {0,0...}
	arr2 := [5]int{1}               // {1, 0...0}
	arr3 := [...]int{1, 2, 3, 4, 5} // arr2 is an array with cap of 5
	fmt.Println("arr: ", arr, ", arr2: ", arr2)
	fmt.Println("arr3: ", arr3)
	s := []int{5} // s is a slice
	s = append(s, 7)
	s = append(s, 9)
	fmt.Printf("len(s) = %d, cap(s)=%d, ptr(s)=%p\n", len(s), cap(s), &s[0])
	x := append(s, 11)
	fmt.Printf("len(s) = %d, cap(s)=%d, ptr(s)=%p\n", len(s), cap(s), &s[0])
	fmt.Printf("len(x) = %d, cap(x)=%d, ptr(x)=%p\n", len(x), cap(x), &x[0])
	y := append(s, 12)
	fmt.Printf("len(s) = %d, cap(s)=%d, ptr(s)=%p\n", len(s), cap(s), &s[0])
	fmt.Printf("len(y) = %d, cap(y)=%d, ptr(y)=%p\n", len(y), cap(y), &y[0])
	fmt.Println(s, x, y)

}
