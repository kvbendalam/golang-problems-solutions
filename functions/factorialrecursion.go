package main

import "fmt"

func main() {
	number := 5
	fmt.Println(factorial(number))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)

}
