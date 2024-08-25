package main

import (
	"fmt"
	"os"

	"github.com/coljac/decomment/internal/processor"
)

/* ONE COMMENT */
const BEFORE = /* ok */ "before" // comment one
// comment two
func main() {
	if len(os.Args) > 1 {
		for _, path := range os.Args[1:] {
			file, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", path, err)
				continue
			}
			defer file.Close()
			processor.Process(file, path)
		}
	} else {
		processor.Process(os.Stdin, "")
	}
}
