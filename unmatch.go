package piscine

func Unmatch(arr []int) int {
	for i, num := range arr {
		found := false

		for j, otherNum := range arr {
			if i != j && num == otherNum {
				found = true
				break
			}
		}

		if !found {
			return num
		}
	}

	return -1
}
