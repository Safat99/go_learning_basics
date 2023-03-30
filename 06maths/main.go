package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to math in golang")

	// var numberOne int = 2
	// var numberTwo float64 = 4.5

	// fmt.Println("The sum is: ", numberOne+int(numberTwo))

	//random number generation
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(3) + 1)

}
