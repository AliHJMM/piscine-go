package piscine

func IsUpper(s string) bool {
	boolean := true
	slice := []rune(s)
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			boolean = true
		} else {
			boolean = false
		}
	}
	return boolean
}
