package graphs

import (
	"reflect"
	"slices"
	"testing"
)

func TestDFSAndBFS(t *testing.T) {
	adj := [][]int{{1, 2}, {3}, {3}, {4}, {}}
	if got := DFS(adj, 0); !reflect.DeepEqual(got, []int{0, 1, 3, 4, 2}) {
		t.Fatalf("dfs got %v", got)
	}
	if got := BFS(adj, 0); !reflect.DeepEqual(got, []int{0, 1, 2, 3, 4}) {
		t.Fatalf("bfs got %v", got)
	}
}

func TestTopologicalSort(t *testing.T) {
	adj := [][]int{{1, 2}, {3}, {3}, {}}
	got, err := TopologicalSort(adj)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, []int{0, 1, 2, 3}) && !reflect.DeepEqual(got, []int{0, 2, 1, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestTopologicalSortCycle(t *testing.T) {
	_, err := TopologicalSort([][]int{{1}, {0}})
	if err == nil {
		t.Fatal("expected cycle error")
	}
}

func TestDijkstra(t *testing.T) {
	matrix := [][]int{{0, 4, 1, 0}, {0, 0, 2, 5}, {0, 0, 0, 1}, {0, 0, 0, 0}}
	result, err := Dijkstra(matrix, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{0, 4, 1, 2}
	if !reflect.DeepEqual(result.Distances, want) {
		t.Fatalf("got %v want %v", result.Distances, want)
	}
}

func TestDijkstraErrors(t *testing.T) {
	if _, err := Dijkstra([][]int{{0}}, -1); err == nil {
		t.Fatal("expected invalid start error")
	}
	if _, err := Dijkstra([][]int{{0, -1}, {0, 0}}, 0); err == nil {
		t.Fatal("expected negative weight error")
	}
}

func TestPrim(t *testing.T) {
	matrix := [][]int{{0, 2, 0, 6, 0}, {2, 0, 3, 8, 5}, {0, 3, 0, 0, 7}, {6, 8, 0, 0, 9}, {0, 5, 7, 9, 0}}
	result, err := Prim(matrix, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.TotalWeight != 16 {
		t.Fatalf("got total %d", result.TotalWeight)
	}
	if len(result.Edges) != 4 {
		t.Fatalf("got %d edges", len(result.Edges))
	}
}

func TestPrimErrors(t *testing.T) {
	if _, err := Prim([][]int{{0}}, -1); err == nil {
		t.Fatal("expected invalid start error")
	}
	if _, err := Prim([][]int{{0, -1}, {0, 0}}, 0); err == nil {
		t.Fatal("expected negative weight error")
	}
}

func TestBellmanFord(t *testing.T) {
	matrix := [][]int{
		{0, 4, 5, 0},
		{0, 0, -2, 6},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	result, err := BellmanFord(matrix, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{0, 4, 2, 3}
	if !reflect.DeepEqual(result.Distances, want) {
		t.Fatalf("got %v want %v", result.Distances, want)
	}
}

func TestBellmanFordNegativeCycle(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0},
		{0, 0, -2},
		{0, -2, 0},
	}
	if _, err := BellmanFord(matrix, 0); err == nil {
		t.Fatal("expected negative cycle error")
	}
}

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(5)
	if uf.Count() != 5 {
		t.Fatalf("count = %d", uf.Count())
	}
	if !uf.Union(0, 1) || !uf.Union(1, 2) {
		t.Fatal("expected unions to merge sets")
	}
	if !uf.Connected(0, 2) {
		t.Fatal("expected nodes 0 and 2 to be connected")
	}
	if uf.Find(2) != uf.Find(0) {
		t.Fatal("expected path compression to preserve representative")
	}
	if uf.Count() != 3 {
		t.Fatalf("count = %d", uf.Count())
	}
}

func TestKruskal(t *testing.T) {
	edges := []Edge{
		{From: 0, To: 1, Weight: 10},
		{From: 0, To: 2, Weight: 6},
		{From: 0, To: 3, Weight: 5},
		{From: 1, To: 3, Weight: 15},
		{From: 2, To: 3, Weight: 4},
	}
	result, err := Kruskal(edges, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.TotalWeight != 19 {
		t.Fatalf("got total %d", result.TotalWeight)
	}
	if len(result.Edges) != 3 {
		t.Fatalf("got %d edges", len(result.Edges))
	}
	if result.Components != 1 {
		t.Fatalf("components = %d", result.Components)
	}
}

func TestKruskalDisconnectedGraph(t *testing.T) {
	edges := []Edge{
		{From: 0, To: 1, Weight: 1},
		{From: 2, To: 3, Weight: 2},
	}
	result, err := Kruskal(edges, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Components != 3 {
		t.Fatalf("components = %d", result.Components)
	}
}

func TestTarjan(t *testing.T) {
	adj := [][]int{
		{1},
		{2, 3},
		{0},
		{4},
		{5, 7},
		{6},
		{4},
		{},
	}
	got := normalizeComponents(Tarjan(adj))
	want := normalizeComponents([][]int{{0, 1, 2}, {3}, {4, 5, 6}, {7}})
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestCycleDetection(t *testing.T) {
	if !HasCycleDirected([][]int{{1}, {2}, {0}}) {
		t.Fatal("expected directed cycle")
	}
	if HasCycleDirected([][]int{{1}, {2}, {}}) {
		t.Fatal("did not expect directed cycle")
	}
	if !HasCycleUndirected([][]int{{1, 2}, {0, 2}, {0, 1}}) {
		t.Fatal("expected undirected cycle")
	}
	if HasCycleUndirected([][]int{{1}, {0, 2}, {1}}) {
		t.Fatal("did not expect undirected cycle")
	}
}

func TestConnectedComponents(t *testing.T) {
	adj := [][]int{
		{1},
		{0},
		{3},
		{2},
		{},
	}
	got := normalizeComponents(ConnectedComponents(adj))
	want := normalizeComponents([][]int{{0, 1}, {2, 3}, {4}})
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func normalizeComponents(components [][]int) [][]int {
	out := make([][]int, len(components))
	for i := range components {
		out[i] = append([]int(nil), components[i]...)
		slices.Sort(out[i])
	}
	slices.SortFunc(out, func(a, b []int) int {
		limit := len(a)
		if len(b) < limit {
			limit = len(b)
		}
		for i := 0; i < limit; i++ {
			if a[i] != b[i] {
				if a[i] < b[i] {
					return -1
				}
				return 1
			}
		}
		switch {
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
		default:
			return 0
		}
	})
	return out
}
