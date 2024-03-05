package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	wordStart := -1

	for i, c := range str {
		isSpace := c == ' ' || c == '\t' || c == '\n' || c == '\r'
		if wordStart == -1 && !isSpace {
			wordStart = i
		} else if wordStart != -1 && (isSpace || i == len(str)-1) {
			if isSpace {
				word := str[wordStart:i]
				if word != "" {
					summary[word]++
				}
			} else {
				summary[str[wordStart:]]++
			}
			wordStart = -1
		}
	}

	return summary
}
