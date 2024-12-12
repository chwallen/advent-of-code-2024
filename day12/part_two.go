package day12

import (
	"io"
)

func PartTwo(r io.Reader) any {
	price := 0
	for region := range getCropRegions(r) {
		price += region.area * region.sides
	}
	return price
}
