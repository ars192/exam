package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
	}
	runes := []rune(args[0])
	for _, letter := range runes {
		if letter >= 65 && letter <= 90 {
			for j := letter - 64; j > 0; j-- {
				z01.PrintRune(letter)
			}
		} else if letter >= 97 && letter <= 122 {
			for gg := letter - 96; gg > 0; gg-- {
				z01.PrintRune(letter)
			}
		}
	}
	z01.PrintRune('\n')
}
