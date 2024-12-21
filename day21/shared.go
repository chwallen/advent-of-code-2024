package day21

import (
	"io"

	"aoc/common"
)

const (
	partOneRobots = 3
	partTwoRobots = 26
)

type pathKey struct {
	start, end rune
}

type pathCacheKey struct {
	path  string
	robot int
}

func calculateComplexitySum(r io.Reader, robots int) int {
	pathCache := make(map[pathCacheKey]int, 1000)
	sum := 0
	for line := range common.ReadLinesLazy(r) {
		numberPart := common.Atoi(line[0 : len(line)-1])
		length := getPathLength(line, robots, pathCache)
		sum += length * numberPart
	}
	return sum
}

func getPathLength(
	path string,
	robot int,
	pathCache map[pathCacheKey]int,
) int {
	cacheKey := pathCacheKey{path, robot}
	if cachedValue, exists := pathCache[cacheKey]; exists {
		return cachedValue
	}

	nextRobot := robot - 1
	length := 0
	current := 'A'
	for _, next := range path {
		// Pressing the same button repeatedly
		if current == next {
			length += 1
		} else {
			p := paths[pathKey{current, next}]
			if nextRobot == 0 {
				length += len(p)
			} else {
				length += getPathLength(p, nextRobot, pathCache)
			}
		}
		current = next
	}

	pathCache[cacheKey] = length
	return length
}

// These are the most efficient paths from one symbol to the next, including
// pressing A.
var paths = map[pathKey]string{
	// Directional paths
	pathKey{'A', '^'}: "<A",
	pathKey{'A', '>'}: "vA",
	pathKey{'A', 'v'}: "<vA",
	pathKey{'A', '<'}: "v<<A",
	pathKey{'^', 'A'}: ">A",
	pathKey{'^', 'v'}: "vA",
	pathKey{'^', '>'}: "v>A",
	pathKey{'^', '<'}: "v<A",
	pathKey{'>', 'A'}: "^A",
	pathKey{'>', '^'}: "<^A",
	pathKey{'>', 'v'}: "<A",
	pathKey{'>', '<'}: "<<A",
	pathKey{'v', 'A'}: "^>A",
	pathKey{'v', '^'}: "^A",
	pathKey{'v', '>'}: ">A",
	pathKey{'v', '<'}: "<A",
	pathKey{'<', 'A'}: ">>^A",
	pathKey{'<', '^'}: ">^A",
	pathKey{'<', '>'}: ">>A",
	pathKey{'<', 'v'}: ">A",
	// Numerical paths
	// A -> x
	pathKey{'A', '0'}: "<A",
	pathKey{'A', '1'}: "^<<A",
	pathKey{'A', '2'}: "<^A",
	pathKey{'A', '3'}: "^A",
	pathKey{'A', '4'}: "^^<<A",
	pathKey{'A', '5'}: "<^^A",
	pathKey{'A', '6'}: "^^A",
	pathKey{'A', '7'}: "^^^<<A",
	pathKey{'A', '8'}: "<^^^A",
	pathKey{'A', '9'}: "^^^A",
	// 0 -> x
	pathKey{'0', 'A'}: ">A",
	pathKey{'0', '1'}: "^<A",
	pathKey{'0', '2'}: "^A",
	pathKey{'0', '3'}: "^>A",
	pathKey{'0', '4'}: "^^<A",
	pathKey{'0', '5'}: "^^A",
	pathKey{'0', '6'}: "^^>A",
	pathKey{'0', '7'}: "^^^<A",
	pathKey{'0', '8'}: "^^^A",
	pathKey{'0', '9'}: "^^^>A",
	// 1 -> x
	pathKey{'1', 'A'}: ">>vA",
	pathKey{'1', '0'}: ">vA",
	pathKey{'1', '2'}: ">A",
	pathKey{'1', '3'}: ">>A",
	pathKey{'1', '4'}: "^A",
	pathKey{'1', '5'}: "^>A",
	pathKey{'1', '6'}: "^>>A",
	pathKey{'1', '7'}: "^^A",
	pathKey{'1', '8'}: "^^>A",
	pathKey{'1', '9'}: "^^>>A",
	// 2 -> x
	pathKey{'2', 'A'}: "v>A",
	pathKey{'2', '0'}: "vA",
	pathKey{'2', '1'}: "<A",
	pathKey{'2', '3'}: ">A",
	pathKey{'2', '4'}: "<^A",
	pathKey{'2', '5'}: "^A",
	pathKey{'2', '6'}: "^>A",
	pathKey{'2', '7'}: "<^^A",
	pathKey{'2', '8'}: "^^A",
	pathKey{'2', '9'}: "^^>A",
	// 3 -> x
	pathKey{'3', 'A'}: "vA",
	pathKey{'3', '0'}: "<vA",
	pathKey{'3', '1'}: "<<A",
	pathKey{'3', '2'}: "<A",
	pathKey{'3', '4'}: "<<^A",
	pathKey{'3', '5'}: "<^A",
	pathKey{'3', '6'}: "^A",
	pathKey{'3', '7'}: "<<^^A",
	pathKey{'3', '8'}: "<^^A",
	pathKey{'3', '9'}: "^^A",
	// 4 -> x
	pathKey{'4', 'A'}: ">>vvA",
	pathKey{'4', '0'}: ">vvA",
	pathKey{'4', '1'}: "vA",
	pathKey{'4', '2'}: "v>A",
	pathKey{'4', '3'}: "v>>A",
	pathKey{'4', '5'}: ">A",
	pathKey{'4', '6'}: ">>A",
	pathKey{'4', '7'}: "^A",
	pathKey{'4', '8'}: "^>A",
	pathKey{'4', '9'}: "^>>A",
	// 5 -> x
	pathKey{'5', 'A'}: "vv>A",
	pathKey{'5', '0'}: "vvA",
	pathKey{'5', '1'}: "<vA",
	pathKey{'5', '2'}: "vA",
	pathKey{'5', '3'}: "v>A",
	pathKey{'5', '4'}: "<A",
	pathKey{'5', '6'}: ">A",
	pathKey{'5', '7'}: "<^A",
	pathKey{'5', '8'}: "^A",
	pathKey{'5', '9'}: "^>A",
	// 6 -> x
	pathKey{'6', 'A'}: "vvA",
	pathKey{'6', '0'}: "<vvA",
	pathKey{'6', '1'}: "<<vA",
	pathKey{'6', '2'}: "<vA",
	pathKey{'6', '3'}: "vA",
	pathKey{'6', '4'}: "<<A",
	pathKey{'6', '5'}: "<A",
	pathKey{'6', '7'}: "<<^A",
	pathKey{'6', '8'}: "<^A",
	pathKey{'6', '9'}: "^A",
	// 7 -> x
	pathKey{'7', 'A'}: ">>vvvA",
	pathKey{'7', '0'}: ">vvvA",
	pathKey{'7', '1'}: "vvA",
	pathKey{'7', '2'}: "vv>A",
	pathKey{'7', '3'}: "vv>>A",
	pathKey{'7', '4'}: "vA",
	pathKey{'7', '5'}: "v>A",
	pathKey{'7', '6'}: "v>>A",
	pathKey{'7', '8'}: ">A",
	pathKey{'7', '9'}: ">>A",
	// 8 -> x
	pathKey{'8', 'A'}: "vvv>A",
	pathKey{'8', '0'}: "vvvA",
	pathKey{'8', '1'}: "<vvA",
	pathKey{'8', '2'}: "vvA",
	pathKey{'8', '3'}: "vv>A",
	pathKey{'8', '4'}: "<vA",
	pathKey{'8', '5'}: "vA",
	pathKey{'8', '6'}: "v>A",
	pathKey{'8', '7'}: "<A",
	pathKey{'8', '9'}: ">A",
	// 9 -> x
	pathKey{'9', 'A'}: "vvvA",
	pathKey{'9', '0'}: "<vvvA",
	pathKey{'9', '1'}: "<<vvA",
	pathKey{'9', '2'}: "<vvA",
	pathKey{'9', '3'}: "vvA",
	pathKey{'9', '4'}: "<<vA",
	pathKey{'9', '5'}: "<vA",
	pathKey{'9', '6'}: "vA",
	pathKey{'9', '7'}: "<<A",
	pathKey{'9', '9'}: "<A",
}
