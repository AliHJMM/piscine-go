package piscine

func Unmatch(arr []int) int {
	for _, num := range arr {
		count := 0

		for _, otherNum := range arr {
			if num == otherNum {
				count++
			}
		}
		// {1, 2, 3, 1, 2, 3, 4}
		if count%2 != 0 {
			return num
		}
	}

	return -1
}
