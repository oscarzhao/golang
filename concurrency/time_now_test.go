package concurrency

import (
	"testing"
	"time"
)

// Benchmark test for time.Now()
// under linux, expected 40-50 ns/op
// under macos, expected 700-800 ns/op
func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

// run with `go test --bench=.`
