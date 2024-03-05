package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			z01.PrintRune(rune('0' + i))
			z01.PrintRune(rune('0' + j))

			if i != 0 || j != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
