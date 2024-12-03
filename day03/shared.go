package day03

import "regexp"

var multiplicationRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
