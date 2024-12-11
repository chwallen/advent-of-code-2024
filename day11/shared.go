package day11

import (
	"io"
	"strconv"
	"strings"

	"aoc/common"
)

const (
	partOneGenerations = 25
	partTwoGenerations = 75
)

func getNumberOfStones(r io.Reader, generations int) int {
	line := common.ReadLinesEager(r)[0]
	currentGeneration := make(map[string]int, 1000)
	nextGeneration := make(map[string]int, 1000)

	for _, inputStone := range strings.Split(line, " ") {
		currentGeneration[inputStone]++
	}

	for i := 0; i < generations; i++ {
		for stoneNumber, count := range currentGeneration {
			if stoneNumber == "0" {
				nextGeneration["1"] += count
				continue
			}

			if len(stoneNumber)%2 == 0 {
				half := len(stoneNumber) / 2
				firstHalf := stoneNumber[:half]
				secondHalf := strings.TrimLeft(stoneNumber[half:], "0")
				if secondHalf == "" {
					secondHalf = "0"
				}
				nextGeneration[firstHalf] += count
				nextGeneration[secondHalf] += count
			} else {
				stoneAsInt := common.Atoi(stoneNumber)
				stone := strconv.Itoa(stoneAsInt * 2024)
				nextGeneration[stone] += count
			}
		}
		currentGeneration, nextGeneration = nextGeneration, currentGeneration
		clear(nextGeneration)
	}

	stones := 0
	for _, count := range currentGeneration {
		stones += count
	}
	return stones
}
