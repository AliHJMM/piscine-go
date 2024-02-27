package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for _, chr := range args[0] {
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')
}
