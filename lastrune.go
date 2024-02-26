package piscine

func LastRune(s string) rune {
	slice := []rune(s)
	return slice[len(slice)-1]
}
