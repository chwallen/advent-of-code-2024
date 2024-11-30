package common

import (
	"fmt"
)

// Check checks if err is not nil and panics if so.
func Check(err error, format string, a ...any) {
	if err != nil {
		message := fmt.Sprintf(format, a...)
		panic(fmt.Errorf("%s: %s", message, err))
	}
}
