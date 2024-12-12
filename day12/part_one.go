package day12

import (
	"io"
)

func PartOne(r io.Reader) any {
	price := 0
	for region := range getCropRegions(r) {
		price += region.area * region.perimeter
	}
	return price
}
