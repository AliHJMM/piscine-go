package piscine

func BasicAtoi(s string) int {
	slice := []rune(s)
	ans := 0
	for i := 0; i < len(s); i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return ans
		} else {
			ans *= 10
			ans += int(slice[i]) - '0'
		}
	}
	return ans
}
