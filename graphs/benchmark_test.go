package graphs

import (
	"fmt"
	"testing"
)

func benchmarkDirectedAdjacency(nodes int, extraStride int) [][]int {
	adj := make([][]int, nodes)
	for from := 0; from < nodes; from++ {
		if from+1 < nodes {
			adj[from] = append(adj[from], from+1)
		}
		if from+extraStride < nodes {
			adj[from] = append(adj[from], from+extraStride)
		}
		if from+extraStride+1 < nodes {
			adj[from] = append(adj[from], from+extraStride+1)
		}
	}
	return adj
}

func benchmarkWeightedMatrix(nodes int, dense bool) [][]int {
	matrix := make([][]int, nodes)
	for i := range matrix {
		matrix[i] = make([]int, nodes)
	}
	for from := 0; from < nodes; from++ {
		for to := from + 1; to < nodes; to++ {
			if !dense && to-from > 3 {
				continue
			}
			weight := (from+to)%9 + 1
			matrix[from][to] = weight
			matrix[to][from] = weight
		}
	}
	return matrix
}

func benchmarkWeightedEdges(nodes int, dense bool) []Edge {
	edges := make([]Edge, 0, nodes*3)
	for from := 0; from < nodes; from++ {
		for to := from + 1; to < nodes; to++ {
			if !dense && to-from > 3 {
				continue
			}
			edges = append(edges, Edge{
				From:   from,
				To:     to,
				Weight: (from+to)%9 + 1,
			})
		}
	}
	return edges
}

func BenchmarkDFS(b *testing.B) {
	for _, nodes := range []int{16, 64, 256, 1024} {
		adj := benchmarkDirectedAdjacency(nodes, 3)
		b.Run(fmt.Sprintf("DFS/n=%d", nodes), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = DFS(adj, 0)
			}
		})
	}
}

func BenchmarkBFS(b *testing.B) {
	for _, nodes := range []int{16, 64, 256, 1024} {
		adj := benchmarkDirectedAdjacency(nodes, 4)
		b.Run(fmt.Sprintf("BFS/n=%d", nodes), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = BFS(adj, 0)
			}
		})
	}
}

func BenchmarkTopologicalSort(b *testing.B) {
	for _, nodes := range []int{16, 64, 256, 1024} {
		adj := benchmarkDirectedAdjacency(nodes, 5)
		b.Run(fmt.Sprintf("TopologicalSort/n=%d", nodes), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = TopologicalSort(adj)
			}
		})
	}
}

func BenchmarkDijkstra(b *testing.B) {
	for _, nodes := range []int{16, 64, 128, 256} {
		for _, dense := range []bool{false, true} {
			matrix := benchmarkWeightedMatrix(nodes, dense)
			shape := "sparse"
			if dense {
				shape = "dense"
			}
			b.Run(fmt.Sprintf("Dijkstra/%s/n=%d", shape, nodes), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = Dijkstra(matrix, 0)
				}
			})
		}
	}
}

func BenchmarkPrim(b *testing.B) {
	for _, nodes := range []int{16, 64, 128, 256} {
		for _, dense := range []bool{false, true} {
			matrix := benchmarkWeightedMatrix(nodes, dense)
			shape := "sparse"
			if dense {
				shape = "dense"
			}
			b.Run(fmt.Sprintf("Prim/%s/n=%d", shape, nodes), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = Prim(matrix, 0)
				}
			})
		}
	}
}

func BenchmarkBellmanFord(b *testing.B) {
	for _, nodes := range []int{16, 64, 128} {
		for _, dense := range []bool{false, true} {
			matrix := benchmarkWeightedMatrix(nodes, dense)
			shape := "sparse"
			if dense {
				shape = "dense"
			}
			b.Run(fmt.Sprintf("BellmanFord/%s/n=%d", shape, nodes), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = BellmanFord(matrix, 0)
				}
			})
		}
	}
}

func BenchmarkKruskal(b *testing.B) {
	for _, nodes := range []int{16, 64, 128, 256} {
		for _, dense := range []bool{false, true} {
			edges := benchmarkWeightedEdges(nodes, dense)
			shape := "sparse"
			if dense {
				shape = "dense"
			}
			b.Run(fmt.Sprintf("Kruskal/%s/n=%d", shape, nodes), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = Kruskal(edges, nodes)
				}
			})
		}
	}
}
