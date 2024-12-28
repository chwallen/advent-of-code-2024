package day23

import (
	"io"
	"slices"
	"strings"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	graph := buildComputerGraph(r)

	R := common.NewSet[string]()
	P := common.NewSet[string]()
	X := common.NewSet[string]()
	for node := range graph {
		P.Add(node)
	}

	var largestClique []string
	findLargestClique(graph, R, P, X, &largestClique)
	slices.Sort(largestClique)

	return strings.Join(largestClique, ",")
}

// Bron-Kerbosch algorithm with pivot
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func findLargestClique(
	graph map[string]common.Set[string],
	R, P, X common.Set[string],
	largestClique *[]string,
) {
	if P.Len() == 0 && X.Len() == 0 {
		if R.Len() > len(*largestClique) {
			*largestClique = R.Items()
		}
		return
	}

	pivotNode := findNodeWithLargestNeighborhood(graph, P.Union(X))
	// Process nodes that aren't in the pivot's neighborhood
	for node := range P.Difference(graph[pivotNode]).All() {
		newR := R.Clone()
		newR.Add(node)
		findLargestClique(
			graph,
			newR,
			P.Intersection(graph[node]),
			X.Intersection(graph[node]),
			largestClique,
		)

		P.Delete(node)
		X.Add(node)
	}
}

func findNodeWithLargestNeighborhood(
	graph map[string]common.Set[string],
	nodes common.Set[string],
) string {
	var node string
	for n := range nodes.All() {
		if graph[n].Len() > graph[node].Len() {
			node = n
		}
	}
	return node
}
