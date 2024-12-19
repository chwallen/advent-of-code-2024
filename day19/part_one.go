package day19

import (
	"io"
	"strings"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)

	rootNode := createTree(lines[0])

	knownDesigns := common.NewSet[string]()
	for _, pattern := range strings.Split(lines[0], ", ") {
		knownDesigns.Add(pattern)
	}

	possibleDesigns := 0
	for _, line := range lines[2:] {
		if isDesignPossible(rootNode, line, knownDesigns) {
			possibleDesigns += 1
		}
	}
	return possibleDesigns
}

func isDesignPossible(root *node, design string, knownDesigns common.Set[string]) bool {
	if knownDesigns.Contains(design) {
		return true
	}

	currentNode := root
	for i := 0; i < len(design); i++ {
		currentNode = currentNode.getChild(design[i])
		if currentNode == nil {
			return false
		}

		patternLength := len(currentNode.getPattern())
		if patternLength > 0 && isDesignPossible(root, design[patternLength:], knownDesigns) {
			knownDesigns.Add(design)
			return true
		}
	}
	return false
}
