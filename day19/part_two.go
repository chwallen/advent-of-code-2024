package day19

import (
	"io"
	"strings"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)

	rootNode := createTree(lines[0])

	knownDesigns := make(map[string]int, 20_000)
	for _, pattern := range strings.Split(lines[0], ", ") {
		countPossibleDesigns(rootNode, pattern, knownDesigns)
	}

	possibleDesignCombinations := 0
	for _, line := range lines[2:] {
		possibleDesignCombinations += countPossibleDesigns(rootNode, line, knownDesigns)
	}
	return possibleDesignCombinations
}

func countPossibleDesigns(root *node, design string, knownDesigns map[string]int) int {
	count, isKnownDesign := knownDesigns[design]
	if isKnownDesign {
		return count
	}
	currentNode := root
	for i := 0; i < len(design); i++ {
		currentNode = currentNode.getChild(design[i])
		if currentNode == nil {
			knownDesigns[design] = count
			return count
		} else if design == currentNode.getPattern() {
			count += 1
			knownDesigns[design] = count
			return count
		}
		patternLength := len(currentNode.getPattern())
		if patternLength > 0 {
			possibleDesigns := countPossibleDesigns(root, design[patternLength:], knownDesigns)
			knownDesigns[design] = possibleDesigns
			count += possibleDesigns
		}
	}
	knownDesigns[design] = count
	return count
}
