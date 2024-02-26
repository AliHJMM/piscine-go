package piscine

func ToLower(s string) string {
	slice := []rune(s)
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] += 32
		}
	}
	return string(slice)
}
