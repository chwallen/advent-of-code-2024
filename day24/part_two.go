package day24

import (
	"io"
	"slices"
	"strings"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	connections, wires := parseGatesAndWires(r)

	lastZ := "z00"
	for s, _ := range wires {
		if s[0] == 'z' {
			n := common.Atoi(s[1:])
			if n > common.Atoi(lastZ[1:]) {
				lastZ = s
			}
		}
	}

	invalidWires := common.NewSet[string]()
	for _, gates := range connections {
		for _, g := range gates {
			if g.isZOutput() && g.op != "XOR" && g.output != lastZ {
				invalidWires.Add(g.output)
			} else if g.op == "XOR" && !g.isZOutput() && !g.isXAndYInputs() {
				invalidWires.Add(g.output)
			} else if g.op == "AND" && g.input1 != "x00" && g.input2 != "x00" {
				for _, subConnection := range connections[g.output] {
					if subConnection.op != "OR" {
						invalidWires.Add(g.output)
					}
				}
			} else if g.op == "XOR" {
				for _, subConnection := range connections[g.output] {
					if subConnection.op == "OR" {
						invalidWires.Add(g.output)
					}
				}
			}
		}
	}

	return strings.Join(slices.Sorted(invalidWires.All()), ",")
}
