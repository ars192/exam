package main

import "fmt"

func main() {
	s := "Hello World"
	arg := []rune(s)
	var revarg []rune
	i := len(arg) - 1
	fmt.Println(i)
	for j := 0; j <= i; j++ {
		revarg[j] = arg[i-j]
		fmt.Println(j)
	}
	//for i >= 0 && string(s[i]) == " " {
	// 	i = i - 1
	// }
	//
	fmt.Println(revarg)
}
