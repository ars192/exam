// "https://public.01-edu.or…jects/sortwordarrprog.en"
// 0	"--cast"
// 1	"make"
package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(array []string) {
	count := 0
	for range array {
		count++
	}
	for i := 0; i < count; i++ {

	}
}
