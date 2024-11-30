package common

import (
	"bufio"
	"io"
	"iter"
	"os"
	"slices"
)

// CloseFile closes the file and panics if it was already closed.
func CloseFile(file *os.File, name string) {
	err := file.Close()
	Check(err, "unable to close file %s", name)
}

// ReadLinesEager eagerly reads all lines from reader. Panics on error.
func ReadLinesEager(r io.Reader) []string {
	return slices.Collect(ReadLinesLazy(r))
}

// ReadLinesLazy lazily reads all lines from reader. Panics on error.
func ReadLinesLazy(r io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		scanner := bufio.NewScanner(r)
		for i := 0; scanner.Scan(); i++ {
			if !yield(scanner.Text()) {
				return
			}
		}
		Check(scanner.Err(), "error reading lines")
	}
}

// ReadFileToString reads the contents of filePath to a string. Panics on error.
func ReadFileToString(filePath string) string {
	f, err := os.Open(filePath)
	Check(err, "unable to open file %s", filePath)
	defer CloseFile(f, filePath)

	bytes, err := io.ReadAll(f)
	Check(err, "unable to read file %s", filePath)

	return string(bytes)
}
