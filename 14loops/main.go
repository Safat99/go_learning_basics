package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops")

	// loops over slice
	days := []string{"Sunday", "Monday", "Tuesday"}
	fmt.Println(days)

	// normal loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println()

	// another way
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println()

	// another way
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// as don't have any while loop

	tempValue := 1

	for tempValue < 10 {

		if tempValue == 3 {
			goto safat
		}

		fmt.Println("value is: ", tempValue)
		tempValue++
	}

	//for the goto statement
safat:
	fmt.Println("jumping at the sarkersafat99@gmai.com ")

}
