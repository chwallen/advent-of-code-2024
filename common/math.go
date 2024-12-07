package common

// Abs gets the absolute value of v.
func Abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

// DivRem divides a with b and returns both the quotient and the remainder.
func DivRem(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a - b*quotient
	return quotient, remainder
}
