package piscine

func AlphaCount(s string) int {
	count := 0
	for _, ch := range s {
		if ch >= 'A' && ch <= 'z' {
			count++
		}
	}
	return count
}
