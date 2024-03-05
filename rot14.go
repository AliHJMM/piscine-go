package piscine

func Rot14(str string) string {
	output := ""
	for _, chr := range str {
		if chr >= 'A' && chr < 'M' || chr >= 'a' && chr < 'm' {
			output += string(chr + 14)
		} else if chr >= 'M' && chr <= 'Z' || chr >= 'm' && chr <= 'z' {
			output += string(chr - 12)
		} else {
			output += string(chr)
		}
	}
	return output
}
