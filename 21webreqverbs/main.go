package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to web verb video - LCO")
	//PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"let's go with Golang",
			"price": 1500,
			"platform" : "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {

	const myurl = "http://localhost:8000/postform"

	// creating formdata

	data := url.Values{}

	data.Add("firstname", "safat")
	data.Add("lastname", "mahmud")
	data.Add("email", "safat@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
