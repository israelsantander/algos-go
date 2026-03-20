package sorting

import (
	"reflect"
	"testing"
)

func intLess(a, b int) bool {
	return a < b
}

func sortingCases() []struct {
	name string
	in   []int
	want []int
} {
	return []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", nil, nil},
		{"single", []int{5}, []int{5}},
		{"duplicates", []int{4, 2, 4, 1, 2}, []int{1, 2, 2, 4, 4}},
		{"sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}
}

func TestSortingAlgorithms(t *testing.T) {
	algorithms := map[string]func([]int) []int{
		"Bubble":    Bubble[int],
		"Counting":  Counting,
		"Selection": Selection[int],
		"Insertion": Insertion[int],
		"Radix":     Radix,
		"Shell":     Shell[int],
		"Merge":     Merge[int],
		"Quick":     Quick[int],
		"Heap":      Heap[int],
	}

	for algoName, fn := range algorithms {
		for _, tt := range sortingCases() {
			t.Run(algoName+"/"+tt.name, func(t *testing.T) {
				input := append([]int(nil), tt.in...)
				got := fn(input)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("got %v want %v", got, tt.want)
				}
				if !reflect.DeepEqual(input, tt.in) {
					t.Fatalf("copying variant mutated input: got %v want %v", input, tt.in)
				}
			})
		}
	}
}

func TestSortingFuncAlgorithms(t *testing.T) {
	algorithms := map[string]func([]int, func(int, int) bool) []int{
		"BubbleFunc":    BubbleFunc[int],
		"SelectionFunc": SelectionFunc[int],
		"InsertionFunc": InsertionFunc[int],
		"ShellFunc":     ShellFunc[int],
		"MergeFunc":     MergeFunc[int],
		"QuickFunc":     QuickFunc[int],
		"HeapFunc":      HeapFunc[int],
	}

	for algoName, fn := range algorithms {
		for _, tt := range sortingCases() {
			t.Run(algoName+"/"+tt.name, func(t *testing.T) {
				input := append([]int(nil), tt.in...)
				got := fn(input, intLess)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("got %v want %v", got, tt.want)
				}
				if !reflect.DeepEqual(input, tt.in) {
					t.Fatalf("copying Func variant mutated input: got %v want %v", input, tt.in)
				}
			})
		}
	}
}

func TestSortingInPlace(t *testing.T) {
	algorithms := map[string]func([]int){
		"BubbleInPlace":    BubbleInPlace[int],
		"CountingInPlace":  CountingInPlace,
		"SelectionInPlace": SelectionInPlace[int],
		"InsertionInPlace": InsertionInPlace[int],
		"RadixInPlace":     RadixInPlace,
		"ShellInPlace":     ShellInPlace[int],
		"MergeInPlace":     MergeInPlace[int],
		"QuickInPlace":     QuickInPlace[int],
		"HeapInPlace":      HeapInPlace[int],
	}

	for name, fn := range algorithms {
		t.Run(name, func(t *testing.T) {
			values := []int{5, 1, 4, 2, 3}
			fn(values)
			if !reflect.DeepEqual(values, []int{1, 2, 3, 4, 5}) {
				t.Fatalf("got %v", values)
			}
		})
	}
}

func TestSortingInPlaceFuncAlgorithms(t *testing.T) {
	algorithms := map[string]func([]int, func(int, int) bool){
		"BubbleInPlaceFunc":    BubbleInPlaceFunc[int],
		"SelectionInPlaceFunc": SelectionInPlaceFunc[int],
		"InsertionInPlaceFunc": InsertionInPlaceFunc[int],
		"ShellInPlaceFunc":     ShellInPlaceFunc[int],
		"MergeInPlaceFunc":     MergeInPlaceFunc[int],
		"QuickInPlaceFunc":     QuickInPlaceFunc[int],
		"HeapInPlaceFunc":      HeapInPlaceFunc[int],
	}

	for name, fn := range algorithms {
		for _, tt := range sortingCases() {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				values := append([]int(nil), tt.in...)
				fn(values, intLess)
				if !reflect.DeepEqual(values, tt.want) {
					t.Fatalf("got %v want %v", values, tt.want)
				}
			})
		}
	}
}

func TestFuncVariantsHandleDuplicatePivotValues(t *testing.T) {
	tests := map[string]func([]int, func(int, int) bool){
		"MergeInPlaceFunc": MergeInPlaceFunc[int],
		"QuickInPlaceFunc": QuickInPlaceFunc[int],
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			values := []int{7, 3, 7, 2, 7, 1, 7, 2, 7}
			fn(values, intLess)
			if !reflect.DeepEqual(values, []int{1, 2, 2, 3, 7, 7, 7, 7, 7}) {
				t.Fatalf("got %v", values)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	if got := Reverse(input); !reflect.DeepEqual(got, []int{4, 3, 2, 1}) {
		t.Fatalf("got %v", got)
	}
	if !reflect.DeepEqual(input, []int{1, 2, 3, 4}) {
		t.Fatalf("reverse mutated input")
	}

	ReverseInPlace(input)
	if !reflect.DeepEqual(input, []int{4, 3, 2, 1}) {
		t.Fatalf("reverse in place got %v", input)
	}
}

func TestIntegerSortsHandleNegatives(t *testing.T) {
	values := []int{12, -5, 7, 0, -5, 3}
	want := []int{-5, -5, 0, 3, 7, 12}
	if got := Counting(values); !reflect.DeepEqual(got, want) {
		t.Fatalf("counting got %v want %v", got, want)
	}
	if got := Radix(values); !reflect.DeepEqual(got, want) {
		t.Fatalf("radix got %v want %v", got, want)
	}
}

func TestQuickSelect(t *testing.T) {
	values := []int{9, 1, 7, 3, 5}
	got, ok := QuickSelect(values, 2)
	if !ok || got != 5 {
		t.Fatalf("got %v %v want 5 true", got, ok)
	}
	if !reflect.DeepEqual(values, []int{9, 1, 7, 3, 5}) {
		t.Fatalf("quickselect mutated input: %v", values)
	}
	if _, ok := QuickSelect(values, -1); ok {
		t.Fatal("expected out of range to fail")
	}
	if _, ok := QuickSelect(values, len(values)); ok {
		t.Fatal("expected out of range to fail")
	}
}

func TestQuickSelectFunc(t *testing.T) {
	type item struct {
		score int
	}
	values := []item{{9}, {1}, {7}, {3}, {5}}
	got, ok := QuickSelectFunc(values, 2, func(a, b item) bool {
		return a.score < b.score
	})
	if !ok || got.score != 5 {
		t.Fatalf("got %v %v want score=5 true", got, ok)
	}
}

func TestFuncVariantsPanicOnNilComparator(t *testing.T) {
	tests := map[string]func(){
		"BubbleFunc":        func() { _ = BubbleFunc([]int{2, 1}, nil) },
		"BubbleInPlaceFunc": func() { BubbleInPlaceFunc([]int{2, 1}, nil) },
		"MergeFunc":         func() { _ = MergeFunc([]int{2, 1}, nil) },
		"QuickSelectFunc":   func() { _, _ = QuickSelectFunc([]int{2, 1}, 0, nil) },
		"QuickInPlaceFunc":  func() { QuickInPlaceFunc([]int{2, 1}, nil) },
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatalf("expected panic")
				}
			}()
			fn()
		})
	}
}
