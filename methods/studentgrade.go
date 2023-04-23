package main

import "fmt"

type Student struct {
	name       string
	rollnumber int
	marks      int
	grade      string
}

func (s Student) calculateGrade() string {
	if s.marks > 90 && s.marks < 100 {
		return "A+"
	} else if s.marks > 80 && s.marks < 89 {
		return "B+"
	} else if s.marks > 70 && s.marks < 79 {
		return "C+"
	} else if s.marks > 60 && s.marks < 69 {
		return "D+"
	} else {
		return "fail"
	}
}

func main() {
	s := Student{"Krishna", 1, 68, ""}

	fmt.Println(s.calculateGrade())
}
