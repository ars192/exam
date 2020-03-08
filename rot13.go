package main

import (
	"github.com/01-edu/z01"
	"os"
	//"fmt"
)

func main() {

	args := os.Args[1:]
	str := args[0]
	runes := []rune(str)
	
	for i, s := range runes {

		if (s < 'a' && s > 'z') || (s < 'A' && s > 'Z') {
			continue
		}
		
		for j := 0; j < 13; j++ {
			
			if s == 'z' {
				runes[i] = 'a' - 1
			}
			if s == 'Z' {
				runes[i] = 'A' - 1
			}
			runes[i]++
		}
		
	}
	for _, s := range runes {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
