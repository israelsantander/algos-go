package lists_test

import (
	"fmt"

	"github.com/israelsantander/algos-go/lists"
)

func ExampleSinglyLinkedList() {
	var list lists.SinglyLinkedList[int]
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	fmt.Println(list.Values())
	// Output: [0 1 2]
}
