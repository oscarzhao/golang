/*************************************************************************
 *  @author:      Programming in go, chapter 1.4 and oscar
 *  @date created:  2015-04-10
 *  @purpose:       1. 2-d slice
 ************************************************************************/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//    digitMap := char[][]{
	//      {' ', ' ', '2', '2', '2', ' ']
	//stringList := [5]string{"1", "2", "3"} //array declared
	//for index, val := range stringList {
	//fmt.Println(index, val)
	//}

	digitMap := [][]string{
		{"  000  ",
			" 0   0 ",
			"0     0",
			"0     0",
			"0     0",
			" 0   0 ",
			"  000  "},
		{" 1 ",
			"11 ",
			" 1 ",
			" 1 ",
			" 1 ",
			" 1 ",
			"111"},
		{"  222  ",
			" 2   2 ",
			"    2  ",
			"   2   ",
			" 2     ",
			"2      ",
			"2222222"},
		{"  333  ",
			" 3   3 ",
			"      3",
			"    33 ",
			"      3",
			" 3   3 ",
			"  333  "},
		{"    4  ",
			"   44  ",
			"  4 4  ",
			" 4  4  ",
			"4444444",
			"    4  ",
			"    4  "},
		{"5555555",
			"5      ",
			"5555   ",
			"    5  ",
			"      5",
			" 5    5",
			"  5555 "},
		{"6666666",
			"6      ",
			"6      ",
			"6666666",
			"6     6",
			"6     6",
			"6666666"},
		{" 777777",
			"      7",
			"     7 ",
			"    7  ",
			"   7   ",
			"  7    ",
			" 7     "},
		{"8888888",
			"8     8",
			"8     8",
			"8888888",
			"8     8",
			"8     8",
			"8888888"},
		{"9999999",
			"9     9",
			"9     9",
			"9999999",
			"      9",
			"      9",
			"      9"},
	}

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	start := time.Now().UnixNano()
	input := os.Args[1]
	i := 0
	for i < len(input) {
		fmt.Printf("%s", "********")
		i++
	}
	fmt.Println()
	i = 0
	for i < 7 {
		for _, c := range input {
			fmt.Printf("%s ", digitMap[c-'0'][i])
		}
		fmt.Println()
		i++
	}
	i = 0
	for i < len(input) {
		fmt.Printf("%s", "********")
		i++
	}
	fmt.Println()
	diff := (time.Now().UnixNano() - start) / 1000000 // ms
	fmt.Printf("execute in %d ms\n", diff)
}
