package main

/*
	An example to use gouuid to generate random uuid string
*/

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
)

func main() {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	hexStr := u4.String()
	fmt.Printf("hex       :%v\nhex string:%v\n", u4, hexStr)
}
