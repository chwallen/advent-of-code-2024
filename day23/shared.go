package day23

import (
	"io"
	"strings"

	"aoc/common"
)

func buildComputerGraph(r io.Reader) map[string]common.Set[string] {
	computerGraph := make(map[string]common.Set[string], 200)
	for line := range common.ReadLinesLazy(r) {
		computers := strings.Split(line, "-")
		a, b := computers[0], computers[1]

		addConnection(computerGraph, a, b)
		addConnection(computerGraph, b, a)
	}

	return computerGraph
}

func addConnection(graph map[string]common.Set[string], nodeA string, nodeB string) {
	if _, setExists := graph[nodeA]; !setExists {
		graph[nodeA] = common.NewSet[string]()
	}
	graph[nodeA].Add(nodeB)
}
