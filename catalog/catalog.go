package catalog

import (
	"github.com/israelsantander/algos-go/graphs"
	"github.com/israelsantander/algos-go/recursion"
	"github.com/israelsantander/algos-go/searching"
	"github.com/israelsantander/algos-go/sorting"
)

type Family string

const (
	FamilySorting   Family = "sorting"
	FamilySearching Family = "searching"
	FamilyGraphs    Family = "graphs"
	FamilyRecursion Family = "recursion"
)

type Entry struct {
	ID         string
	Name       string
	Family     Family
	Complexity string
	Example    string
}

// Entries returns the catalog metadata for the algorithms exposed by this module.
func Entries() []Entry {
	return []Entry{
		{ID: "bubble", Name: "Bubble Sort", Family: FamilySorting, Complexity: "O(n^2)", Example: "sorting.Bubble([]int{5,1,4,2,8})"},
		{ID: "counting", Name: "Counting Sort", Family: FamilySorting, Complexity: "O(n+k)", Example: "sorting.Counting([]int{4,-1,2,-1})"},
		{ID: "quickselect", Name: "QuickSelect", Family: FamilySorting, Complexity: "O(n) average", Example: "sorting.QuickSelect([]int{9,1,7,3,5}, 2)"},
		{ID: "selection", Name: "Selection Sort", Family: FamilySorting, Complexity: "O(n^2)", Example: "sorting.Selection([]int{64,25,12,22,11})"},
		{ID: "insertion", Name: "Insertion Sort", Family: FamilySorting, Complexity: "O(n^2)", Example: "sorting.Insertion([]int{31,41,59,26})"},
		{ID: "binary", Name: "Binary Search", Family: FamilySearching, Complexity: "O(log n)", Example: "searching.Binary([]int{1,3,5,7,9}, 7)"},
		{ID: "exponential", Name: "Exponential Search", Family: FamilySearching, Complexity: "O(log n)", Example: "searching.Exponential([]int{1,3,5,7,9}, 7)"},
		{ID: "avl", Name: "AVL Tree Search", Family: FamilySearching, Complexity: "O(log n)", Example: "searching.AVLSearch(root, 7)"},
		{ID: "dfs", Name: "Depth-First Search", Family: FamilyGraphs, Complexity: "O(V+E)", Example: "graphs.DFS(adj, 0)"},
		{ID: "dijkstra", Name: "Dijkstra", Family: FamilyGraphs, Complexity: "O(V^2)", Example: "graphs.Dijkstra(matrix, 0)"},
		{ID: "bellman-ford", Name: "Bellman-Ford", Family: FamilyGraphs, Complexity: "O(VE)", Example: "graphs.BellmanFord(matrix, 0)"},
		{ID: "kruskal", Name: "Kruskal", Family: FamilyGraphs, Complexity: "O(E log E)", Example: "graphs.Kruskal(edges, 4)"},
		{ID: "tarjan", Name: "Tarjan SCC", Family: FamilyGraphs, Complexity: "O(V+E)", Example: "graphs.Tarjan(adj)"},
		{ID: "nqueens", Name: "N-Queens", Family: FamilyRecursion, Complexity: "Backtracking", Example: "recursion.NQueens(4)"},
	}
}

var _ = sorting.Bubble[int]
var _ = searching.Binary[int]
var _ = graphs.DFS
var _ = recursion.NQueens
