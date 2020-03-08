package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	arg := os.Args[1]
	runes := []rune(arg)

	if len(arr) == 1 {
		for _, i := range runes {
			if i >= 'A' && i <= 'Z' {
				z01.PrintRune('Z' - (i - 'A'))
			} else if i >= 'a' && i <= 'z' {
				z01.PrintRune('z' - (i - 'a'))
			} else {
				z01.PrintRune(i)
			}
		}
	}
	z01.PrintRune('\n')
}
