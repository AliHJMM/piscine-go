package piscine

func NRune(s string, n int) rune {
	slice := []rune(s)
	return slice[n-1]
}
