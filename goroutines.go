//go routines

package main

import (
	"fmt"
	"time"
)

func main() {
	go computeSquare(5) // for computation it is taking time
	go computeKube(5)   // for computation it is taking time

	time.Sleep(time.Second)
	//takes some time to compute. but while checking this tasks are not completed
}

func computeSquare(num int) {
	result := num * num
	fmt.Println(result)
}

func computeKube(num int) {
	result := num * num * num
	fmt.Println(result)
}

//for go routines we need to use "go" keyword

//if we go infront of any of the function then it will act as go routine
