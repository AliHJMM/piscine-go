package main

import "fmt"

func BasicAtoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		digit := s[i]
		result = result*10 + int(digit-'0')
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
