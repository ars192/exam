// "https://public.01-edu.org/subjects/tabmult.en"
// 0	"--cast"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"os.*"
// 3	"len"
// 4	"strconv.Atoi"
package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		return
	}
	num, err := strconv.Atoi(arguments[0])
	if err != nil {
		return
	}
	if num <= 0 {
		return
	}
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + 48))
		z01.PrintRune(32)
		z01.PrintRune('x')
		z01.PrintRune(32)
		for _, letter := range printNbr(num) {
			z01.PrintRune(letter)
		}
		z01.PrintRune(32)
		z01.PrintRune('=')
		z01.PrintRune(32)
		for _, letter := range printNbr(i * num) {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

func printNbr(num int) string {
	res := ""
	if num/10 != 0 {
		res = res + printNbr(num/10)
	}
	r := '0'
	for i := 0; i < num%10; i++ {
		r++
	}
	return res + string(r)
}
