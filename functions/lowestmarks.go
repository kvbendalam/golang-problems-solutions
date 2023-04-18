package main

import "fmt"

func main() {
	marks := map[string]int{
		"Alice":  99,
		"Bob":    100,
		"Charle": 45,
		"Darwin": 98,
	}

	lowestMarks := marks["Alice"]
	var studentName string

	for name, marks := range marks {
		if marks < lowestMarks {
			lowestMarks = marks
			studentName = name
		}
	}
	fmt.Println(studentName, lowestMarks)

}