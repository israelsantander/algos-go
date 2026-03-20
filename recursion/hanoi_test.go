package recursion

import (
	"reflect"
	"testing"
)

func TestHanoi(t *testing.T) {
	moves := Hanoi(3, "A", "B", "C")
	if len(moves) != 7 {
		t.Fatalf("got %d moves", len(moves))
	}
	if moves[0].From != "A" || moves[0].To != "C" {
		t.Fatalf("first move = %+v", moves[0])
	}
	if moves[len(moves)-1].From != "A" || moves[len(moves)-1].To != "C" {
		t.Fatalf("last move = %+v", moves[len(moves)-1])
	}
}

func TestFactorial(t *testing.T) {
	if got := Factorial(-1); got != 0 {
		t.Fatalf("got %d want 0", got)
	}
	if got := Factorial(0); got != 1 {
		t.Fatalf("got %d want 1", got)
	}
	if got := Factorial(5); got != 120 {
		t.Fatalf("got %d want 120", got)
	}
}

func TestFibonacci(t *testing.T) {
	if got := Fibonacci(-1); got != 0 {
		t.Fatalf("got %d want 0", got)
	}
	if got := Fibonacci(0); got != 0 {
		t.Fatalf("got %d want 0", got)
	}
	if got := Fibonacci(7); got != 13 {
		t.Fatalf("got %d want 13", got)
	}
}

func TestPermutations(t *testing.T) {
	got := Permutations([]int{1, 2, 3})
	if len(got) != 6 {
		t.Fatalf("got %d permutations", len(got))
	}
	if !reflect.DeepEqual(got[0], []int{1, 2, 3}) {
		t.Fatalf("unexpected first permutation: %v", got[0])
	}
}

func TestCombinations(t *testing.T) {
	got := Combinations([]int{1, 2, 3, 4}, 2)
	want := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestSubsets(t *testing.T) {
	got := Subsets([]int{1, 2})
	want := [][]int{{}, {2}, {1}, {1, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestNQueens(t *testing.T) {
	if got := NQueens(1); !reflect.DeepEqual(got, [][]int{{0}}) {
		t.Fatalf("got %v", got)
	}
	if got := NQueens(2); len(got) != 0 {
		t.Fatalf("expected no solutions, got %v", got)
	}
	if got := NQueens(4); len(got) != 2 {
		t.Fatalf("expected 2 solutions, got %v", got)
	}
}
