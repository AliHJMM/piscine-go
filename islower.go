package piscine

func IsLower(s string) bool {
	boolean := false
	slice := []rune(s)
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 'a' && slice[i] <= 'z' {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
