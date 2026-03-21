package {{package}}

import (
	"reflect"
	"testing"
)

func Test{{Algorithm}}(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.in...)
			got := {{Algorithm}}(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		})
	}
}
