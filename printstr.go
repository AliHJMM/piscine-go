package piscine

import "fmt"

func PrintStr(s string) {
	arr := []rune(s)
	fmt.Print(string(arr))
}
