package main

import "fmt"

func main(){
    var arr[5]int
    arr[0] = 10
    arr[4] = 100
    var i int = 0
    fmt.Println(arr)
    for i < 5 {
        fmt.Println(arr[i])
        i++
    }
}