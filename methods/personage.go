package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) IsAdult() bool {
	if p.age > 18 {
		return true
	}
	return false
}

func main() {
	p := Person{"Krishna", 28}
	fmt.Println(p.IsAdult())
}
