package main

import (
	"fmt"
	"strconv"
)

// https://en.wikipedia.org/wiki/Decimal64_floating-point_format
// this example is to verify why float64 support precision 16

func main() {
	v52 := "1111111111111111111111111111111111111111111111111111" // 52, is precise
	if s, err := strconv.ParseInt(v52, 2, 64); err == nil {
		fmt.Printf("%T, %v,len(s):%d, %v\n", s, s, len(fmt.Sprintf("%d", s)), float64(s))
	}

	v54 := "111111111111111111111111111111111111111111111111111111" // 54, lose precision
	if s, err := strconv.ParseInt(v54, 2, 64); err == nil {
		fmt.Printf("%T, %v,len(s):%d, %v\n", s, s, len(fmt.Sprintf("%d", s)), float64(s))
	}

	v56 := "11111111111111111111111111111111111111111111111111111111" // 56, lose precision
	if s, err := strconv.ParseInt(v56, 2, 64); err == nil {
		fmt.Printf("%T, %v,len(s):%d, %v\n", s, s, len(fmt.Sprintf("%d", s)), float64(s))
	}
}

/* output:
int64, 4503599627370495,len(s):16, 4.503599627370495e+15
int64, 18014398509481983,len(s):17, 1.8014398509481984e+16
int64, 72057594037927935,len(s):17, 7.205759403792794e+16
*/
