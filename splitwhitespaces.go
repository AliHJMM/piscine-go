package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	var currentWord string

	for _, chr := range s {
		if chr == ' ' || chr == '\t' || chr == '\n' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(chr)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
