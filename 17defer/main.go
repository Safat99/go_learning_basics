package main

import "fmt"

func main() {
	defer fmt.Println("joy bangla")
	defer fmt.Println("Hello world")
	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
