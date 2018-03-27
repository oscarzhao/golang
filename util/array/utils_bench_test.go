package array

import (
	"testing"
)

func BenchmarkToInterface(b *testing.B) {
	in := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		_ = ToInterfaces(in)
	}
}

// run with `go test --bench=.`
