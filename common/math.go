package common

// Abs gets the absolute value of v.
func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}
