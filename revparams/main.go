package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := len(args) - 1; i > 0; i-- {
		if i > 0 {
			for _, chr := range args[i] {
				z01.PrintRune(chr)
			}
		}
		z01.PrintRune('\n')
	}
}
