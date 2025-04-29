package bench

import (
	"fmt"
	"time"
)

type BenchResult struct {
	Name       string
	Iterations int
	Min        time.Duration
	Max        time.Duration
	Avg        time.Duration
}

func (br BenchResult) String() string {
	return fmt.Sprintf("Benchmark Results '%s' (%d iterations):\n    min: %s, max: %s, avg: %s",
		br.Name, br.Iterations, br.Min, br.Max, br.Avg)
}

func Run(name string, f func(), n int) BenchResult {
	result := BenchResult{
		Name:       name,
		Iterations: n,
		Min:        time.Duration(1<<63 - 1),
		Max:        time.Duration(0),
		Avg:        time.Duration(0),
	}

	for range n {
		start := time.Now()
		f()
		duration := time.Since(start)

		if duration < result.Min {
			result.Min = duration
		}
		if duration > result.Max {
			result.Max = duration
		}

		result.Avg += duration
	}

	result.Avg /= time.Duration(n)
	return result
}
