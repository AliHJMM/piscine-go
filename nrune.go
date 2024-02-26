package piscine

func NRune(s string, n int) rune {
	slice := []rune(s)
	length := 0
	for i := range slice {
		length = i
	}

	if n-1 >= 0 && n-1 <= length {
		return slice[n-1]
	}
	return 0
}
