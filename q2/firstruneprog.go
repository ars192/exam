package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	for _, k := range s {
		z01.PrintRune(k)
		break
	}

}
