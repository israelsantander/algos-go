package linear_test

import (
	"fmt"

	"github.com/israelsantander/algos-go/linear"
)

func ExampleStack() {
	var s linear.Stack[int]
	s.Push(10)
	s.Push(20)
	v, _ := s.Pop()
	fmt.Println(v)
	// Output: 20
}
