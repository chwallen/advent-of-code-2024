package day23

import (
	"io"
	"slices"
	"strings"
)

func PartOne(r io.Reader) any {
	graph := buildComputerGraph(r)

	result := 0
	uniqueConnections := make(map[string]bool)

	for nodeA, neighborsSet := range graph {
		neighbours := neighborsSet.Items()
		for i := 0; i < neighborsSet.Len(); i++ {
			for j := i + 1; j < neighborsSet.Len(); j++ {
				nodeB, nodeC := neighbours[i], neighbours[j]
				if (nodeA[0] == 't' || nodeB[0] == 't' || nodeC[0] == 't') && graph[nodeB].Contains(nodeC) {
					nodes := []string{nodeA, nodeB, nodeC}
					slices.Sort(nodes)
					key := strings.Join(nodes, ",")
					if !uniqueConnections[key] {
						uniqueConnections[key] = true
						result += 1
					}
				}
			}
		}
	}

	return result
}
