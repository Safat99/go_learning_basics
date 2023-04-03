package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// it uses the library which is much more powerful
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content) //we have to write this line

	fmt.Println("ByteCOunt is: ", byteCount)

	// whatever the data it is holding inside, it will convert into String
	fmt.Println(responseString.String())

	fmt.Println(string(content)) // as a beginner we can use this

}
