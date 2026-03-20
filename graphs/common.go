package graphs

import (
	"errors"
	"math"
)

var (
	errInvalidStart       = errors.New("invalid start node")
	errNegativeWeights    = errors.New("negative weights are not supported")
	errNegativeCycle      = errors.New("graph contains a negative-weight cycle")
	errGraphContainsCycle = errors.New("graph contains a cycle")
	errInvalidNodeCount   = errors.New("invalid node count")
	errDisconnectedGraph  = errors.New("graph is disconnected")
)

const infDistance = math.MaxInt / 4

// validNode reports whether node can index a graph with size nodes.
func validNode(node, nodes int) bool {
	return node >= 0 && node < nodes
}

// normalizeUndirectedEdge orders endpoints so undirected duplicates compare consistently.
func normalizeUndirectedEdge(edge Edge) Edge {
	if edge.To < edge.From {
		edge.From, edge.To = edge.To, edge.From
	}
	return edge
}
