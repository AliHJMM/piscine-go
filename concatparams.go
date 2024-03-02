package piscine

func ConcatParams(args []string) string {
	output := ""

	for i, chr := range args {
		output += string(chr)
		if i != (len(args) - 1) {
			output += "\n"
		}
	}
	return output
}
