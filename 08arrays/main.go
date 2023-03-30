package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in golang")

	// array declaration
	var nums = [5]int{1, 2, 3, 4}
	fmt.Println("first array default: ", nums)
	fmt.Printf("type of the nums is: %T\n", nums)
	fmt.Println("the len that is registered: ", len(nums))
	fmt.Println()

	words := [5]string{}
	fmt.Println("first array default: ", words)
	fmt.Printf("type of the words is: %T\n", words)
	fmt.Println("the len that is registered: ", len(words))

	//2d array declaration
	var twoD [2][3]int
	var i int
	var j int
	for i = 0; i < 2; i++ {
		for j = 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue
			}
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD array default: ", twoD)

}
