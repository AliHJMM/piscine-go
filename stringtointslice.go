package piscine

func StringToIntSlice(str string) []int {
	var intSlice []int

	for _, char := range str {
		intValue := int(char)
		intSlice = append(intSlice, intValue)
	}

	return intSlice
}
