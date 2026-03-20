package recursion_test

import (
	"fmt"

	"github.com/israelsantander/algos-go/recursion"
)

func ExampleHanoi() {
	moves := recursion.Hanoi(2, "A", "B", "C")
	fmt.Println(len(moves))
	// Output: 3
}

func ExamplePermutations() {
	fmt.Println(recursion.Permutations([]int{1, 2}))
	// Output: [[1 2] [2 1]]
}
