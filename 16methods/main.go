/*
functions and methods are almost the same. When a function goes inside a cloud it is called as
a method. As golang don't have classes, the function that goes inside the struct will be called
the metods
*/
package main

import "fmt"

func main() {

	fmt.Println("-------------Structs in Go----------")
	// no inheritence in golang; no super or parent

	safat := User{Name: "sarker safat mahmud", Age: 20, Email: "sarkersafat99@gmail.com"}
	fmt.Println(safat)
	fmt.Printf("safat details are: %+v\n", safat) // %v means value.. %+v used in struct for full details
	fmt.Printf("Name is %v and email is %v.\n", safat.Name, safat.Email)

	safat.GetStatus()
	safat.NewMail()
	fmt.Printf("safat details are: %+v\n", safat) // %v means value.. %+v used in struct for full details

}

type User struct {
	// Camelcase means they are public
	Name   string
	Email  string
	Status bool
	Age    int
}

// if we are planning to export this method the return type has to be Camelcase otherwise will make it lowercase.
func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

// similar to getter-setter
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
