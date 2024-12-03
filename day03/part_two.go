package day03

import (
	"io"
	"strings"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	line := strings.Join(common.ReadLinesEager(r), "")
	sum := 0
	for _, s := range strings.Split(line, "do()") {
		disabledRegionStart := strings.Index(s, "don't()")
		if disabledRegionStart == -1 {
			disabledRegionStart = len(s)
		}
		enabledRegion := s[0:disabledRegionStart]
		for _, match := range multiplicationRegex.FindAllStringSubmatch(enabledRegion, -1) {
			sum += common.Atoi(match[1]) * common.Atoi(match[2])
		}
	}
	return sum
}
