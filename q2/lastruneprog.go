package piscine
}
func LastRune(s string) rune {
	len := 0
	for range s {
		len++
	}
	for i, a := range s {
		if len-1 != i {

		} else {
			z01.PrintRune(a)
		}
	}
}
