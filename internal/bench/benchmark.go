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

type B struct {
	result  *BenchResult
	results map[string]*BenchResult
}

func (b B) String() string {
	str := fmt.Sprintf("Benchmark Result '%s':\n", b.result.Name)
	str += fmt.Sprintf("  (total): %s:", b.result)
	for name, result := range b.results {
		str += fmt.Sprintf("\n  (%s): %s", name, result)
	}

	return str
}

func (b *B) Run(name string, f func(*B)) {
	if _, ok := b.results[name]; !ok {
		b.results[name] = &BenchResult{
			Name:       name,
			Iterations: 0,
			Min:        time.Duration(1<<63 - 1),
			Max:        time.Duration(0),
			Avg:        time.Duration(0),
		}
	}

	t := time.Now()
	f(b)
	d := time.Since(t)

	if d < b.results[name].Min {
		b.results[name].Min = d
	}
	if d > b.results[name].Max {
		b.results[name].Max = d
	}

	b.results[name].Iterations++
	b.results[name].Avg += d
}

func (br BenchResult) String() string {
	return fmt.Sprintf("min: %s, max: %s, avg: %s",
		br.Min, br.Max, br.Avg)
}

func Run(name string, f func(*B), n int) *B {
	result := &BenchResult{
		Name:       name,
		Iterations: n,
		Min:        time.Duration(1<<63 - 1),
		Max:        time.Duration(0),
		Avg:        time.Duration(0),
	}

	b := &B{
		result:  result,
		results: make(map[string]*BenchResult),
	}

	for range n {
		start := time.Now()
		f(b)
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
	for _, result := range b.results {
		result.Avg /= time.Duration(result.Iterations)
	}

	return b
}
