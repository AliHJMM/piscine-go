package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := range args {
		if i > 0 {
			for _, chr := range args[i] {
				z01.PrintRune(chr)
			}
		}
		z01.PrintRune('\n')
	}
}
