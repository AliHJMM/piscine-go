package piscine

import "fmt"
func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb > 20 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	fmt.Print("\n")
	return result
}