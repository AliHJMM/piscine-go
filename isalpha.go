package piscine

func IsAlpha(s string) bool {
	boolean := false
	slice := []rune(s)
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 'a' && slice[i] <= 'z' || slice[i] >= 'A' && slice[i] <= 'Z' || slice[i] >= '0' && slice[i] <= '9' {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
