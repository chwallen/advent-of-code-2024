package day24

import (
	"fmt"
	"io"
	"strings"

	"aoc/common"
)

type gate struct {
	input1 string
	input2 string
	op     string
	output string
}

func (g gate) isXAndYInputs() bool {
	return (g.input1[0] == 'x' && g.input2[0] == 'y') ||
		(g.input1[0] == 'y' && g.input2[0] == 'x')
}

func (g gate) isZOutput() bool {
	return g.output[0] == 'z'
}

func parseGatesAndWires(r io.Reader) (connections map[string][]gate, wires map[string]int) {
	connections = make(map[string][]gate)
	wires = make(map[string]int)

	unprocessedGates := make([]gate, 0, 1000)
	for line := range common.ReadLinesLazy(r) {
		if len(line) == 0 {
			continue
		}
		if line[3] == ':' {
			name, value, _ := strings.Cut(line, ": ")
			wires[name] = common.Atoi(value)
		} else {
			parts := strings.Split(line, " ")
			g := gate{input1: parts[0], input2: parts[2], op: parts[1], output: parts[4]}
			connections[g.input1] = append(connections[g.input1], g)
			connections[g.input2] = append(connections[g.input2], g)

			if !processGate(g, wires) {
				unprocessedGates = append(unprocessedGates, g)
			}
		}
	}

	var g gate
	for len(unprocessedGates) > 0 {
		g, unprocessedGates = unprocessedGates[0], unprocessedGates[1:]
		if !processGate(g, wires) {
			unprocessedGates = append(unprocessedGates, g)
		}
	}

	return connections, wires
}

func processGate(g gate, wires map[string]int) bool {
	left, isLeftOk := wires[g.input1]
	right, isRightOk := wires[g.input2]
	if isLeftOk && isRightOk {
		switch g.op {
		case "AND":
			wires[g.output] = left & right
		case "OR":
			wires[g.output] = left | right
		case "XOR":
			wires[g.output] = left ^ right
		default:
			panic(fmt.Errorf("unknown op %s", g.op))
		}
		return true
	} else {
		return false
	}
}
