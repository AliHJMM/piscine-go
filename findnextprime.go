package piscine

func FindNextPrime(nb int) int {
	for a := nb + 1; ; a++ {
		if IsPrime(a) {
			return a
		}
	}
}
