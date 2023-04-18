// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	marks := []int{99, 98, 97, 23, 23}

	total := total(marks)
	marksLength := len(marks)
	fmt.Println("Total marks are", total)
	fmt.Println("Average marks are", average(total, marksLength))
}

func total(marks []int) int {
	total := 0
	for _, v := range marks {
		total += v
	}

	return total
}

func average(total int, marksLength int) int {
	return total / marksLength
}
