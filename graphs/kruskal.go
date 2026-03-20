package graphs

import "sort"

// Kruskal computes a minimum spanning forest from a weighted edge list.
//
// It returns an error when nodeCount is negative or an edge endpoint is out of range.
func Kruskal(edges []Edge, nodeCount int) (KruskalResult, error) {
	if nodeCount < 0 {
		return KruskalResult{}, errInvalidNodeCount
	}
	if nodeCount == 0 {
		return KruskalResult{}, nil
	}

	normalized := make([]Edge, 0, len(edges))
	for _, edge := range edges {
		if !validNode(edge.From, nodeCount) || !validNode(edge.To, nodeCount) {
			return KruskalResult{}, errInvalidStart
		}
		if edge.From == edge.To {
			continue
		}
		normalized = append(normalized, normalizeUndirectedEdge(edge))
	}

	sort.Slice(normalized, func(i, j int) bool {
		if normalized[i].Weight != normalized[j].Weight {
			return normalized[i].Weight < normalized[j].Weight
		}
		if normalized[i].From != normalized[j].From {
			return normalized[i].From < normalized[j].From
		}
		return normalized[i].To < normalized[j].To
	})

	uf := NewUnionFind(nodeCount)
	result := KruskalResult{
		Edges:      make([]Edge, 0, max(0, nodeCount-1)),
		Components: nodeCount,
	}
	for _, edge := range normalized {
		if uf.Union(edge.From, edge.To) {
			result.Edges = append(result.Edges, edge)
			result.TotalWeight += edge.Weight
			result.Components--
		}
	}
	return result, nil
}
