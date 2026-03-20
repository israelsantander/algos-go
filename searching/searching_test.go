package searching

import (
	"reflect"
	"testing"
)

type searchItem struct {
	ID   int
	Name string
}

func itemCompare(a, b searchItem) int {
	switch {
	case a.ID < b.ID:
		return -1
	case a.ID > b.ID:
		return 1
	default:
		return 0
	}
}

func TestLinear(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		target int
		want   int
	}{
		{"nil", nil, 1, -1},
		{"empty", []int{}, 1, -1},
		{"single hit", []int{5}, 5, 0},
		{"single miss", []int{5}, 8, -1},
		{"first duplicate", []int{7, 3, 9, 3}, 3, 1},
		{"missing", []int{1, 2, 3}, 8, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Linear(tt.values, tt.target); got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestLinearFunc(t *testing.T) {
	values := []searchItem{
		{ID: 2, Name: "b"},
		{ID: 4, Name: "d"},
		{ID: 4, Name: "duplicate"},
		{ID: 7, Name: "g"},
	}

	if got := LinearFunc(values, searchItem{ID: 4}, itemCompare); got != 1 {
		t.Fatalf("got %d want %d", got, 1)
	}
	if got := LinearFunc(values, searchItem{ID: 9}, itemCompare); got != -1 {
		t.Fatalf("got %d want %d", got, -1)
	}
}

func TestBinary(t *testing.T) {
	values := []int{1, 3, 4, 7, 9, 12}
	if got := Binary(values, 7); got != 3 {
		t.Fatalf("got %d want %d", got, 3)
	}
	if got := Binary(values, 2); got != -1 {
		t.Fatalf("got %d want %d", got, -1)
	}
}

func TestBinaryReturnsFirstDuplicate(t *testing.T) {
	values := []int{1, 3, 3, 3, 5, 8}
	if got := Binary(values, 3); got != 1 {
		t.Fatalf("got %d want %d", got, 1)
	}
}

func TestBinaryFunc(t *testing.T) {
	values := []searchItem{
		{ID: 1, Name: "a"},
		{ID: 3, Name: "c"},
		{ID: 3, Name: "c2"},
		{ID: 5, Name: "e"},
	}

	if got := BinaryFunc(values, searchItem{ID: 3}, itemCompare); got != 1 {
		t.Fatalf("got %d want %d", got, 1)
	}
	if got := BinaryFunc(values, searchItem{ID: 4}, itemCompare); got != -1 {
		t.Fatalf("got %d want %d", got, -1)
	}
}

func TestExponential(t *testing.T) {
	values := []int{1, 3, 3, 3, 5, 8, 13}
	if got := Exponential(values, 3); got != 1 {
		t.Fatalf("got %d want %d", got, 1)
	}
	if got := Exponential(values, 9); got != -1 {
		t.Fatalf("got %d want -1", got)
	}
}

func TestJump(t *testing.T) {
	values := []int{1, 3, 3, 3, 5, 8, 13}
	if got := Jump(values, 3); got != 1 {
		t.Fatalf("got %d want %d", got, 1)
	}
	if got := Jump(values, 9); got != -1 {
		t.Fatalf("got %d want -1", got)
	}
}

func TestSortedForBinary(t *testing.T) {
	input := []int{9, 1, 4, 1}
	got := SortedForBinary(input)
	want := []int{1, 1, 4, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
	if !reflect.DeepEqual(input, []int{9, 1, 4, 1}) {
		t.Fatalf("SortedForBinary mutated input: got %v", input)
	}
}

func TestSortedForBinaryFunc(t *testing.T) {
	input := []searchItem{
		{ID: 4, Name: "d"},
		{ID: 1, Name: "a"},
		{ID: 3, Name: "c"},
	}
	got := SortedForBinaryFunc(input, itemCompare)
	want := []searchItem{
		{ID: 1, Name: "a"},
		{ID: 3, Name: "c"},
		{ID: 4, Name: "d"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
	if !reflect.DeepEqual(input, []searchItem{
		{ID: 4, Name: "d"},
		{ID: 1, Name: "a"},
		{ID: 3, Name: "c"},
	}) {
		t.Fatalf("SortedForBinaryFunc mutated input: got %v", input)
	}
}

func TestBST(t *testing.T) {
	root := BuildBST([]int{8, 3, 10, 1, 6, 14, 4, 7, 13})
	if node := SearchBST(root, 7); node == nil || node.Value != 7 {
		t.Fatalf("expected to find 7")
	}
	if node := SearchBST(root, 99); node != nil {
		t.Fatalf("expected nil for missing value")
	}
	want := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	if got := InOrder(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestBSTDuplicateInsertionGoesRight(t *testing.T) {
	root := BuildBST([]int{5, 5, 5})
	if root == nil || root.Right == nil || root.Right.Right == nil {
		t.Fatalf("expected duplicates to be inserted on the right")
	}
	if got := InOrder(root); !reflect.DeepEqual(got, []int{5, 5, 5}) {
		t.Fatalf("got %v", got)
	}
}

func TestBSTFunc(t *testing.T) {
	values := []searchItem{
		{ID: 8, Name: "h"},
		{ID: 3, Name: "c"},
		{ID: 10, Name: "j"},
		{ID: 6, Name: "f"},
	}

	root := BuildBSTFunc(values, itemCompare)
	if node := SearchBSTFunc(root, searchItem{ID: 6}, itemCompare); node == nil || node.Value.ID != 6 {
		t.Fatalf("expected to find item 6")
	}
	if node := SearchBSTFunc(root, searchItem{ID: 99}, itemCompare); node != nil {
		t.Fatalf("expected nil for missing item")
	}
	if got := InOrder(root); !reflect.DeepEqual(got, []searchItem{
		{ID: 3, Name: "c"},
		{ID: 6, Name: "f"},
		{ID: 8, Name: "h"},
		{ID: 10, Name: "j"},
	}) {
		t.Fatalf("got %v", got)
	}
}

func TestAVL(t *testing.T) {
	var root *AVLNode[int]
	for _, value := range []int{1, 2, 3, 4, 5, 6, 7} {
		root = AVLInsert(root, value)
	}
	if root == nil || root.Value != 4 {
		t.Fatalf("unexpected root: %+v", root)
	}
	if node := AVLSearch(root, 6); node == nil || node.Value != 6 {
		t.Fatalf("expected to find 6")
	}
	if got := AVLInOrder(root); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Fatalf("got %v", got)
	}
	if absInt(avlBalance(root)) > 1 {
		t.Fatalf("root is not balanced: %d", avlBalance(root))
	}
}

func TestAVLFunc(t *testing.T) {
	values := []searchItem{
		{ID: 8, Name: "h"},
		{ID: 3, Name: "c"},
		{ID: 10, Name: "j"},
		{ID: 6, Name: "f"},
	}
	var root *AVLNode[searchItem]
	for _, value := range values {
		root = AVLInsertFunc(root, value, itemCompare)
	}
	if node := AVLSearchFunc(root, searchItem{ID: 6}, itemCompare); node == nil || node.Value.ID != 6 {
		t.Fatalf("expected to find item 6")
	}
	if got := AVLInOrder(root); !reflect.DeepEqual(got, []searchItem{
		{ID: 3, Name: "c"},
		{ID: 6, Name: "f"},
		{ID: 8, Name: "h"},
		{ID: 10, Name: "j"},
	}) {
		t.Fatalf("got %v", got)
	}
}

func TestFuncVariantsPanicOnNilComparator(t *testing.T) {
	tests := map[string]func(){
		"LinearFunc":          func() { _ = LinearFunc([]searchItem{{ID: 1}}, searchItem{ID: 1}, nil) },
		"BinaryFunc":          func() { _ = BinaryFunc([]searchItem{{ID: 1}}, searchItem{ID: 1}, nil) },
		"ExponentialFunc":     func() { _ = ExponentialFunc([]searchItem{{ID: 1}}, searchItem{ID: 1}, nil) },
		"JumpFunc":            func() { _ = JumpFunc([]searchItem{{ID: 1}}, searchItem{ID: 1}, nil) },
		"SortedForBinaryFunc": func() { _ = SortedForBinaryFunc([]searchItem{{ID: 1}}, nil) },
		"InsertFunc":          func() { _ = InsertFunc[*searchItem](nil, nil, nil) },
		"BuildBSTFunc":        func() { _ = BuildBSTFunc([]searchItem{{ID: 1}}, nil) },
		"SearchBSTFunc":       func() { _ = SearchBSTFunc((*BSTNode[searchItem])(nil), searchItem{ID: 1}, nil) },
		"AVLInsertFunc":       func() { _ = AVLInsertFunc((*AVLNode[searchItem])(nil), searchItem{ID: 1}, nil) },
		"AVLSearchFunc":       func() { _ = AVLSearchFunc((*AVLNode[searchItem])(nil), searchItem{ID: 1}, nil) },
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

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
