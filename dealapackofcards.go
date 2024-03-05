package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	playersNumber := 4
	cards := len(deck) / playersNumber

	for i := 0; i < playersNumber; i++ {
		start := i * cards
		end := (i + 1) * cards
		playerCards := deck[start:end]

		fmt.Printf("Player %d: ", i+1)
		for _, card := range playerCards {
			printIntAsRunes(card)
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func printIntAsRunes(num int) {
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune

	for num > 0 {
		digit := num % 10
		digits = append(digits, rune(digit+'0'))
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
