package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int // i am a pointer and i will store integer value
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber // creating a pointer and also referencing to some memory
	// refernece means &

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)

}
