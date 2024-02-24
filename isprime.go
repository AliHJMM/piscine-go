package piscine

func IsPrime(nb int) bool {
	rep := true
	if nb == 2 || nb == 3 || nb == 5 || nb == 7 || nb == 11 || nb == 13 || nb == 17 || nb == 19 || nb == 21 || nb == 23 || nb == 29 || nb == 31 || nb == 37 {
		rep = true
	} else if nb == 1 || nb == 0 {
		rep = false
	} else {
		for i := 2; i < 40; i++ {
			if nb%i == 0 {
				rep = false
			}
		}
	}
	return rep
}
