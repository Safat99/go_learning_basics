package main

import "fmt"

const LoginToken string = "ghajksfj" // PUblic

func main() {
	var username string = "safat"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloat float64 = 12555.4556098
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	// default value
	var sampleString string
	fmt.Println(sampleString)
	fmt.Printf("variable is of type : %T \n", sampleString)

	// implicit type of declaring a variable
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 3000
	fmt.Println("number", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

}
