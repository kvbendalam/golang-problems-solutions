package main

import "fmt"

type Employee struct {
	Name        string
	age         int
	Salary      string
	Bloodgroup  string
	Designation string
	Weight      float64
}

func main() {

	// Using new key word - Intializing the struct
	emp := new(Employee)

	emp.Name = "Krishna"
	emp.age = 28
	emp.Salary = "100000"
	emp.Bloodgroup = "B+Ve"
	emp.Designation = "Tech lead"
	emp.Weight = 60.45

	fmt.Println(emp)

	//Directly using Employee Struct

	emp1 := Employee{"Vamsi", 28, "200000", "B-ve", "Tech Manager", 90.45}
	fmt.Println(emp1)

	// Using employee struct with keys as well
	var emp2 = Employee{
		Name:        "Rama",
		age:         30,
		Salary:      "300000",
		Bloodgroup:  "O+ve",
		Designation: "Senior Software Enginerr",
		Weight:      50.0,
	}

	fmt.Println(emp2)
}
