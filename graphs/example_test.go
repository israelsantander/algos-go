package graphs_test

import (
	"fmt"

	"github.com/israelsantander/algos-go/graphs"
)

func ExampleDFS() {
	adj := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(graphs.DFS(adj, 0))
	// Output: [0 1 3 2]
}

func ExampleKruskal() {
	edges := []graphs.Edge{
		{From: 0, To: 1, Weight: 10},
		{From: 0, To: 2, Weight: 6},
		{From: 0, To: 3, Weight: 5},
		{From: 2, To: 3, Weight: 4},
	}

	result, _ := graphs.Kruskal(edges, 4)
	fmt.Println(result.TotalWeight)
	// Output: 19
}
