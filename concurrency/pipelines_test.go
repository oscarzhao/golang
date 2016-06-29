package concurrency

import "testing"

/*
  reference: https://blog.golang.org/pipelines
*/
func TestSimplePipelines(t *testing.T) {
	shouldSuccess := []struct {
		input  []int
		output []int
	}{
		{input: []int{2, 3, 4, 5, 6, 7}, output: []int{4, 9, 16, 25, 36, 49}},
	}
	for _, tc := range shouldSuccess {
		// Set up the pipeline.
		c := gen(tc.input...)
		out := sq(c)

		// Consume the output.
		i := 0
		for o := range out {
			if o != tc.output[i] {
				t.Errorf("input: %d, output should be %d, but is %d\n", tc.input[i], tc.output[i], o)
			}
			i++
		}
	}
}

func TestComplexPipelines(t *testing.T) {
	shouldSuccess := []struct {
		input  []int
		output []int
	}{
		{input: []int{2, 3, 4}, output: []int{16, 81, 256}},
	}
	for _, tc := range shouldSuccess {
		// Set up the pipeline.
		c := gen(tc.input...)
		out := sq(sq(c))

		// Consume the output.
		i := 0
		for o := range out {
			if o != tc.output[i] {
				t.Errorf("input: %d, output should be %d, but is %d\n", tc.input[i], tc.output[i], o)
			}
			i++
		}
	}
}
