package piscine

func IsPrintable(s string) bool {
	boolean := false
	slice := []rune(s)
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 32 && slice[i] <= 126 {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
