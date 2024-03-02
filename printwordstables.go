package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, words := range a {
		output := words
		for _, letters := range output {
			z01.PrintRune(letters)
		}
		z01.PrintRune('\n')
	}
}
