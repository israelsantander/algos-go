package searching_test

import (
	"fmt"

	"github.com/israelsantander/algos-go/searching"
)

func ExampleBinary() {
	fmt.Println(searching.Binary([]int{1, 3, 5, 7, 9}, 7))
	// Output: 3
}

func ExampleBinaryFunc() {
	type person struct {
		ID   int
		Name string
	}

	people := []person{
		{ID: 1, Name: "Ana"},
		{ID: 3, Name: "Bo"},
		{ID: 5, Name: "Eve"},
	}

	target := person{ID: 3}
	index := searching.BinaryFunc(people, target, func(a, b person) int {
		switch {
		case a.ID < b.ID:
			return -1
		case a.ID > b.ID:
			return 1
		default:
			return 0
		}
	})

	fmt.Println(index)
	// Output: 1
}

func ExampleExponential() {
	fmt.Println(searching.Exponential([]int{1, 3, 5, 7, 9}, 7))
	// Output: 3
}
