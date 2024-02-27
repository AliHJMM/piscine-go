package piscine

func BasicJoin(elems []string) string {
	output := " "
	for _, chr := range elems {
		output += chr
	}
	return output
}
