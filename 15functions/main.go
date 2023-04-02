package main

import "fmt"

func main() {
	fmt.Println("Welcome to go functions")
	greet("safat")

	var i int
	var j int
	fmt.Print("Type the first Number: ")
	fmt.Scan(&i)
	fmt.Print("Type the second Number: ")
	fmt.Scan(&j)

	fmt.Println("result for adder is ", adder(i, j))

	// calling the pro Adder
	proResult, _ := proAdder(2, 3, 4, 5)
	fmt.Println("result of the pro adder is:", proResult)

	// already have multiple args in a slice
	nums := []int{1, 2, 3, 4}
	fmt.Println(proAdder(nums...)) // this is the syntax

}

func greet(name string) {
	fmt.Println("hello Mr.", name)
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

// i don't know how many int values are coming
// values will be a slices
// ... is called the variadic functions. They can accept any values
// we can return two values from a func
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "message sent from proAdder"
}
