package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment_id=lkjlkj23"

func main() {
	fmt.Println("welcome to handling URLS in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{ //always need to give the &
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=safat",
	}

	constructedUrl := partsOfUrl.String()
	fmt.Println(constructedUrl)

}
