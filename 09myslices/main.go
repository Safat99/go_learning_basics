package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice tutorial")

	var fruitList = []string{"apple", "mango"}
	fmt.Printf("Type of fruitList is:%T\n", fruitList)

	fruitList = append(fruitList, "peach", "banana")
	fmt.Println(fruitList)
	fmt.Println("The lenghth of fruitList is: ", len(fruitList))

	// fruitList = append(fruitList[1:])
	fmt.Println(fruitList[1:]) //slicing and the last range is non inclusive

	//remove an element from slice
	index := 2
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println("After removing the 2nd index, the fruitList will be: ", fruitList)

	//more on slices -->> memory management
	highScores := make([]int, 4)
	fmt.Printf("type of slice using make() is: %T\n\n", highScores)

	highScores[0] = 235
	highScores[1] = 100
	highScores[2] = 900
	highScores[3] = 800

	highScores = append(highScores, 200, 300)
	fmt.Println("len of highscore slice is now: ", len(highScores))

	//more features on slices -->> sorting which arrays don't have
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
