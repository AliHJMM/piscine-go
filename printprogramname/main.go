package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	slice := []rune(args[0])

	for _, chr := range slice {
		if chr != '.' && chr != '/' {
			z01.PrintRune(chr)
		}
	}
	z01.PrintRune('\n')
}
