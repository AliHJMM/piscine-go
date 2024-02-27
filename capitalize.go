package piscine

func Alpha(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	slice := []rune(s)

	boolean := true

	for i := 0; i < len(s); i++ {
		if Alpha(slice[i]) == true && boolean {
			if slice[i] >= 'a' && slice[i] <= 'z' {
				slice[i] = 'A' - 'a' + slice[i]
			}
			boolean = false
		} else if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] = 'a' - 'A' + slice[i]
		} else if Alpha(slice[i]) == false {
			boolean = true
		}
	}
	return string(slice)
}
