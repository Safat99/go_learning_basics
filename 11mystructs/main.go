package main

import "fmt"

func main() {

	fmt.Println("-------------Structs in Go----------")
	// no inheritence in golang; no super or parent

	safat := User{Name: "sarker safat mahmud", Age: 20, Email: "sarkersafat99@gmail.com"}
	fmt.Println(safat)
	fmt.Printf("safat details are: %+v\n", safat) // %v means value.. %+v used in struct for full details
	fmt.Printf("Name is %v and email is %v.", safat.Name, safat.Email)

}

type User struct {
	// Camelcase means they are public
	Name   string
	Email  string
	Status bool
	Age    int
}
