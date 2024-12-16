package day16

import (
	"container/heap"
	"io"
	"math"

	"aoc/common"
)

type pathStep struct {
	point    common.Point
	previous *pathStep
}

type path struct {
	pos   common.Point
	dir   common.Direction
	score int
	steps *pathStep
}

type mazeTile struct {
	char       rune
	bestScores [4]int
}

func createMaze(r io.Reader) (maze common.Grid[*mazeTile], start common.Point) {
	lines := common.ReadLinesEager(r)
	maze = common.NewGrid[*mazeTile](len(lines[0]), len(lines))

	for y, line := range lines {
		for x, char := range line {
			maze.Set(x, y, &mazeTile{char, [4]int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}})
			if char == 'S' {
				start = common.NewPoint(x, y)
			}
		}
	}
	return maze, start
}

func traverseMaze(r io.Reader) (int, int) {
	maze, start := createMaze(r)
	bestPathsPoints := common.NewSet[common.Point]()
	queue := common.NewPriorityQueue[path](5000, func(a, b path) bool {
		return a.score < b.score
	})

	heap.Push(queue, path{pos: start, dir: common.Right})
	bestScore := math.MaxInt

	for queue.Len() > 0 {
		item := heap.Pop(queue).(path)
		if item.score > bestScore {
			// Since we have a priority queue, we know we have found all good paths
			// when this occurs.
			break
		}

		tile := maze.GetPoint(item.pos)
		if tile.char == 'E' {
			bestScore = item.score
			bestPathsPoints.Add(start)
			for step := item.steps; step != nil; step = step.previous {
				bestPathsPoints.Add(step.point)
			}
		}

		cardinalIndex := item.dir.GetCardinalIndex()
		if tile.bestScores[cardinalIndex] >= item.score {
			tile.bestScores[cardinalIndex] = item.score
			for _, nextDir := range []common.Direction{item.dir, item.dir.TurnRight(), item.dir.TurnLeft()} {
				next := item.pos.Add(nextDir)
				nextScore := item.score + 1
				if nextDir != item.dir {
					nextScore += 1000
				}

				cardinalIndex = nextDir.GetCardinalIndex()
				nextTile := maze.GetPoint(next)
				if nextTile.char != '#' && nextTile.bestScores[cardinalIndex] > nextScore {
					steps := &pathStep{point: next, previous: item.steps}
					heap.Push(queue, path{next, nextDir, nextScore, steps})
				}
			}
		}
	}

	return bestScore, bestPathsPoints.Len()
}
