package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("LCO web request")

	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	fmt.Printf("response is of type: %T/n", response)

	defer response.Body.Close() //caller's responsibility to close this connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
