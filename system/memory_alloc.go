package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

var (
	mem string

	pattern = regexp.MustCompile(`(\d+)([K|k|KB|kb|M|m|MB|mb|G|g|GB|gb])`)
)

func init() {
	flag.StringVar(&mem, "mem", "10M", "assign how much memory to allocate, such 1k, 1m, 1g")
	flag.Parse()
}

func main() {
	arr := pattern.FindAllStringSubmatch(mem, 1)
	if len(arr) != 1 {
		fmt.Printf("invalid param: %s, exit\n", mem)
		return
	}
	if len(arr[0]) != 3 {
		fmt.Printf("Invalid param: %s, exit\n", mem)
		return
	}
	initialSize, err := strconv.Atoi(arr[0][1])
	if err != nil {
		fmt.Printf("invalid param: %s, error:%s\nExit\n", mem, err)
		return
	}
	initialUnit := arr[0][2]

	var totalSizeBytes int64

	switch initialUnit {
	case "KB", "kb", "K", "k":
		totalSizeBytes = 1024 * int64(initialSize)
	case "MB", "mb", "M", "m":
		totalSizeBytes = 1024 * 1024 * int64(initialSize)
	case "GB", "gb", "G", "g":
		totalSizeBytes = 1024 * 1024 * 1024 * int64(initialSize)
	default:
		fmt.Printf("invald unit: %s\nExit\n", initialUnit)
		return
	}
	fmt.Printf("target memory size used: %d%s\n", initialSize, initialUnit)

	var intArr []int64 // 8 bytes

	kbArr := make([]int64, 1024/8, 1024/8)
	for i := int64(0); i < totalSizeBytes/1024; i++ {
		intArr = append(intArr, kbArr...)
	}
	fmt.Printf("len(arr): %d\n", len(intArr))
	for {
	}
}
