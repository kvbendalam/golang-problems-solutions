package main

import "fmt"

type Book struct {
	title  string
	author string
	price  int
}

func (b Book) discount(percentage int) int {
	res := b.price * percentage / 100
	actualPrice := b.price - res
	return actualPrice
}

func main() {
	b := Book{"Introduction to golang", "Dennis", 1000}
	fmt.Println(b.discount(10))
}
