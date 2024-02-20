package main

import "github.com/01-edu/z01"

func printcomb2() {
    for i := 0; i <= 98; i++ {
        for j := i + 1; j <= 99; j++ {
            if i != 0 || j != 1 {
                z01.PrintRune(',')
                z01.PrintRune(' ')
            }
            z01.PrintRune(rune(i/10 + '0'))
            z01.PrintRune(rune(i%10 + '0'))
            z01.PrintRune(' ')
            z01.PrintRune(rune(j/10 + '0'))
            z01.PrintRune(rune(j%10 + '0'))
        }
    }
    z01.PrintRune('$')
    z01.PrintRune('\n')
}

func main() {
    printcomb2()
}
