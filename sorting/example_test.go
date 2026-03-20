package sorting_test

import (
	"fmt"

	"github.com/israelsantander/algos-go/sorting"
)

func ExampleBubble() {
	fmt.Println(sorting.Bubble([]int{5, 1, 4, 2, 8}))
	// Output: [1 2 4 5 8]
}

func ExampleQuickFunc() {
	type person struct {
		Name string
		Age  int
	}

	people := []person{
		{Name: "Ana", Age: 29},
		{Name: "Bo", Age: 18},
		{Name: "Eve", Age: 24},
	}

	sorted := sorting.QuickFunc(people, func(a, b person) bool {
		return a.Age < b.Age
	})

	for _, person := range sorted {
		fmt.Printf("%s:%d ", person.Name, person.Age)
	}

	// Output: Bo:18 Eve:24 Ana:29
}

func ExampleQuickSelect() {
	value, _ := sorting.QuickSelect([]int{9, 1, 7, 3, 5}, 2)
	fmt.Println(value)
	// Output: 5
}
