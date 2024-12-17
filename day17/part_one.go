package day17

import (
	"io"
	"strings"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)

	output := parseIntoProgram(lines).execute()

	var sb strings.Builder
	sb.WriteRune(rune(output[0] + '0'))
	for _, v := range output[1:] {
		sb.WriteRune(',')
		sb.WriteRune(rune(v + '0'))
	}
	return sb.String()
}
