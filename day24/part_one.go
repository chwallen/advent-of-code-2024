package day24

import (
	"io"
	"slices"
)

func PartOne(r io.Reader) any {
	_, wires := parseGatesAndWires(r)

	zWires := make([]string, 0, 50)
	for name := range wires {
		if name[0] == 'z' {
			zWires = append(zWires, name)
		}
	}

	slices.Sort(zWires)
	slices.Reverse(zWires)

	value := 0
	for _, zWire := range zWires {
		bit := wires[zWire]
		value = (value << 1) | bit
	}
	return value
}
