package piscine

func Atoi(s string) int {
	slice := []rune(s)
	ans := 0
	sign := 1
	if slice[0] == '-' {
		sign = -1
		slice = slice[1:]
	} else if slice[0] == '+' {
		slice = slice[1:]
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += int(slice[i]) - '0'
		}
	}
	return ans * sign
}
