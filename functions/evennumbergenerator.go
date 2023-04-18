package main

import "fmt"

func main() {
	fmt.Println(evenNumbers(100))
}

func evenNumbers(x int) []int {
	even := []int{}
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			even = append(even, i)
		}
	}
	return even
}
