package main

import "fmt"

func main() {
	number := 5
	fact := 1
	for i := 1; i <= number; i++ {
		fact = fact * i
	}

	fmt.Println(fact)
}
