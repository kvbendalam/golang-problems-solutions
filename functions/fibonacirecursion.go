package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(0)
	for i := 1; i <= 10; i = i + 1 {
		fmt.Println(Fibonacci(i))
	}
}
