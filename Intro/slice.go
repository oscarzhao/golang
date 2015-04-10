/*************************************************************************
 *  @author:  http://www.zhihu.com/question/27161493
 *  @date created:  2015-04-10
 *  @purpose:  1. strange behavior of slice
 ************************************************************************/
package main
import {
    "fmt"
}

func main(){
    s := []int{5}
    s = append(s, 7)
    s = append(s, 9)
    fmt.Println("s = %s, ptr(s)=%d", s, &s[0])
    x := append(s, 11)
    fmt.Println("s = %s, ptr(s)=%d", s, &s[0])
    fmt.Println("x = %s, ptr(x)=%d", x, &x[0])
    y := append(s, 12)
    fmt.Println("s = %s, ptr(s)=%d", s, &s[0])
    fmt.Println("y = %s, ptr(y)=%d", y, &y[0])
    fmt.Println(s, x, y)

}
