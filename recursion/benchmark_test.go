package recursion

import (
	"fmt"
	"testing"
)

func BenchmarkHanoi(b *testing.B) {
	for _, discs := range []int{3, 5, 8, 10} {
		b.Run(fmt.Sprintf("Hanoi/discs=%d", discs), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = Hanoi(discs, "A", "B", "C")
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	for _, n := range []int{5, 10, 12} {
		b.Run(fmt.Sprintf("Factorial/n=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = Factorial(n)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for _, n := range []int{10, 20, 30} {
		b.Run(fmt.Sprintf("Fibonacci/n=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = Fibonacci(n)
			}
		})
	}
}

func BenchmarkPermutations(b *testing.B) {
	values := []int{1, 2, 3, 4}
	b.Run("Permutations/n=4", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Permutations(values)
		}
	})
}

func BenchmarkNQueens(b *testing.B) {
	for _, n := range []int{4, 6, 8} {
		b.Run(fmt.Sprintf("NQueens/n=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = NQueens(n)
			}
		})
	}
}
