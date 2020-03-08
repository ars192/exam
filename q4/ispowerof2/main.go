package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	lenArgs := 0
	for range args {
		lenArgs++
	}
	if lenArgs != 1 {
		return
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	status := "false"
	for i := 2; i <= num; i = i * 2 {
		if i == num {
			status = "true"
		}
	}
	for _, letter := range status {
		z01.PrintRune(letter)
	}
}
