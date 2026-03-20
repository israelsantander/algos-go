package searching

import (
	"fmt"
	"testing"
)

type benchmarkSearchItem struct {
	ID int
}

func benchmarkItemCompare(a, b benchmarkSearchItem) int {
	switch {
	case a.ID < b.ID:
		return -1
	case a.ID > b.ID:
		return 1
	default:
		return 0
	}
}

func benchmarkSortedInts(size int) []int {
	values := make([]int, size)
	for i := range values {
		values[i] = i * 2
	}
	return values
}

func BenchmarkLinear(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		values := benchmarkSortedInts(size)
		targets := map[string]int{
			"first":   values[0],
			"middle":  values[size/2],
			"last":    values[size-1],
			"missing": -1,
		}
		for label, target := range targets {
			b.Run(fmt.Sprintf("Linear/%s/n=%d", label, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = Linear(values, target)
				}
			})
		}
	}
}

func BenchmarkBinary(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		values := benchmarkSortedInts(size)
		targets := map[string]int{
			"first":   values[0],
			"middle":  values[size/2],
			"last":    values[size-1],
			"missing": -1,
		}
		for label, target := range targets {
			b.Run(fmt.Sprintf("Binary/%s/n=%d", label, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = Binary(values, target)
				}
			})
		}
	}
}

func BenchmarkExponential(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		values := benchmarkSortedInts(size)
		targets := map[string]int{
			"first":   values[0],
			"middle":  values[size/2],
			"last":    values[size-1],
			"missing": -1,
		}
		for label, target := range targets {
			b.Run(fmt.Sprintf("Exponential/%s/n=%d", label, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = Exponential(values, target)
				}
			})
		}
	}
}

func BenchmarkJump(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		values := benchmarkSortedInts(size)
		targets := map[string]int{
			"first":   values[0],
			"middle":  values[size/2],
			"last":    values[size-1],
			"missing": -1,
		}
		for label, target := range targets {
			b.Run(fmt.Sprintf("Jump/%s/n=%d", label, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = Jump(values, target)
				}
			})
		}
	}
}

func BenchmarkBinaryFunc(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		values := make([]benchmarkSearchItem, size)
		for i := range values {
			values[i] = benchmarkSearchItem{ID: i * 2}
		}
		targets := map[string]benchmarkSearchItem{
			"first":   values[0],
			"middle":  values[size/2],
			"last":    values[size-1],
			"missing": {ID: -1},
		}
		for label, target := range targets {
			b.Run(fmt.Sprintf("BinaryFunc/%s/n=%d", label, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = BinaryFunc(values, target, benchmarkItemCompare)
				}
			})
		}
	}
}

func BenchmarkBuildBST(b *testing.B) {
	for _, size := range []int{32, 256, 1024, 4096} {
		values := benchmarkSortedInts(size)
		b.Run(fmt.Sprintf("BuildBST/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = BuildBST(values)
			}
		})
	}
}

func BenchmarkSearchBST(b *testing.B) {
	for _, size := range []int{32, 256, 1024, 4096} {
		values := benchmarkSortedInts(size)
		root := BuildBST(values)
		target := values[size/2]
		b.Run(fmt.Sprintf("SearchBST/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = SearchBST(root, target)
			}
		})
	}
}

func BenchmarkAVLSearch(b *testing.B) {
	for _, size := range []int{32, 256, 1024, 4096} {
		values := benchmarkSortedInts(size)
		var root *AVLNode[int]
		for _, value := range values {
			root = AVLInsert(root, value)
		}
		target := values[size/2]
		b.Run(fmt.Sprintf("AVLSearch/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = AVLSearch(root, target)
			}
		})
	}
}
