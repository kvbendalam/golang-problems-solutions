package main

import "fmt"

func main() {
	marks := []int{98, 99, 100, 22, 33, 32}
	sum, average, highest := f(marks)
	fmt.Println("Sum is", sum, "Average is", average, "Highest is", highest)
}

func f(marks []int) (int, int, int) {
	sum := 0
	for _, v := range marks {
		sum += v
	}

	highest := marks[0]
	for _, v := range marks {
		if v > highest {
			highest = v
		}
	}

	return sum, sum / len(marks), highest
}
