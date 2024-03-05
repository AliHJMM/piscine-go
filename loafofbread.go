package piscine


func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	for i := 0; i < len(str); i += 5 {
		endIndex := i + 5
		if endIndex > len(str) {
			endIndex = len(str)
		}
		result += str[i:endIndex]
		if endIndex < len(str) {
			result += " " + string(str[endIndex])
		}
	}

	return result + "\n"
}

