package {{package}}

import (
	"fmt"
	"math/rand"
	"testing"
)

func Benchmark{{Algorithm}}(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	patterns := []string{"random", "sorted", "reversed", "nearly-sorted"}
	
	for _, size := range sizes {
		for _, pattern := range patterns {
			base := generateData(size, pattern)
			b.Run(fmt.Sprintf("%s/n=%d", pattern, size), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					values := append([]int(nil), base...)
					{{Algorithm}}(values)
				}
			})
		}
	}
}

func generateData(size int, pattern string) []int {
	rng := rand.New(rand.NewSource(42))
	values := make([]int, size)
	switch pattern {
	case "sorted":
		for i := range values {
			values[i] = i
		}
	case "reversed":
		for i := range values {
			values[i] = size - i
		}
	case "nearly-sorted":
		for i := range values {
			values[i] = i
		}
		for i := 0; i < size/20; i++ {
			a, b := rng.Intn(size), rng.Intn(size)
			values[a], values[b] = values[b], values[a]
		}
	default:
		for i := range values {
			values[i] = rng.Intn(size * 10)
		}
	}
	return values
}
