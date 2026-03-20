package sorting

import (
	"fmt"
	"math/rand"
	"testing"
)

func benchmarkSortingAlgorithm(b *testing.B, name string, fn func([]int), sizes []int, patterns []string) {
	for _, size := range sizes {
		for _, pattern := range patterns {
			base := benchmarkInts(size, pattern)
			b.Run(fmt.Sprintf("%s/%s/n=%d", name, pattern, size), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					values := append([]int(nil), base...)
					fn(values)
				}
			})
		}
	}
}

func benchmarkInts(size int, pattern string) []int {
	rng := rand.New(rand.NewSource(int64(size)*7919 + int64(len(pattern))*101))
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
		swaps := max(1, size/20)
		for i := 0; i < swaps; i++ {
			a := rng.Intn(size)
			b := rng.Intn(size)
			values[a], values[b] = values[b], values[a]
		}
	default:
		for i := range values {
			values[i] = rng.Intn(size * 4)
		}
	}
	return values
}

func BenchmarkBubbleInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "BubbleInPlace", BubbleInPlace[int], []int{16, 64, 256}, []string{"random", "sorted", "reversed", "nearly-sorted"})
}

func BenchmarkSelectionInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "SelectionInPlace", SelectionInPlace[int], []int{16, 64, 256}, []string{"random", "sorted", "reversed"})
}

func BenchmarkCountingInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "CountingInPlace", CountingInPlace, []int{32, 128, 512, 4096}, []string{"random", "sorted", "reversed"})
}

func BenchmarkInsertionInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "InsertionInPlace", InsertionInPlace[int], []int{16, 64, 256}, []string{"random", "sorted", "reversed", "nearly-sorted"})
}

func BenchmarkShellInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "ShellInPlace", ShellInPlace[int], []int{32, 128, 512, 2048}, []string{"random", "sorted", "reversed", "nearly-sorted"})
}

func BenchmarkMergeInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "MergeInPlace", MergeInPlace[int], []int{32, 128, 512, 4096}, []string{"random", "sorted", "reversed", "nearly-sorted"})
}

func BenchmarkRadixInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "RadixInPlace", RadixInPlace, []int{32, 128, 512, 4096}, []string{"random", "sorted", "reversed"})
}

func BenchmarkQuickInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "QuickInPlace", QuickInPlace[int], []int{32, 128, 512, 4096}, []string{"random", "sorted", "reversed", "nearly-sorted"})
}

func BenchmarkHeapInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "HeapInPlace", HeapInPlace[int], []int{32, 128, 512, 4096}, []string{"random", "sorted", "reversed"})
}

func BenchmarkReverseInPlace(b *testing.B) {
	benchmarkSortingAlgorithm(b, "ReverseInPlace", ReverseInPlace[int], []int{32, 128, 512, 4096, 16384}, []string{"random", "sorted", "reversed"})
}

func BenchmarkQuickSelect(b *testing.B) {
	for _, size := range []int{32, 128, 512, 4096} {
		base := benchmarkInts(size, "random")
		k := size / 2
		b.Run(fmt.Sprintf("QuickSelect/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = QuickSelect(base, k)
			}
		})
	}
}
