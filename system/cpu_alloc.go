package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	percent float64
)

func init() {
	flag.Float64Var(&percent, "percent", 0.8, "cpu usage by percentage, value must be [0~1], e.g. 0.5 means 50% uage")
	flag.Parse()
}
func main() {
	if percent < 0 || percent > 1 {
		fmt.Printf("percentage must be between 0 and 1.0\n")
		os.Exit(1)
	}
	np := runtime.NumCPU()
	real := int(float64(np) * percent)
	fmt.Printf("cpu core number: %d\n", np)
	fmt.Printf("target cpu percent is adjusted to: %.3f\n", float64(real)/float64(np))

	runtime.GOMAXPROCS(real)
	for i := 0; i < real; i++ {
		go func() {
			for {
			}
		}()

	}
	select {}
}
