package piscine

import (
	"fmt"
)

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	for i := 0; i <= len(a)-1; i++ {
		for _, ch := range a[i] {
			fmt.Print(ch)
		}
}
}

	
