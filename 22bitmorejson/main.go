package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string `json:"website"`
	// - simply means I don't want this field to be reflected
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs bootcamp,", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern bootcamp,", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular bootcamp,", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "ReactJs bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
	`)

	// golang is checking whether the datas are in the correct json format

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // we don't want to send copies
		// that's why we use reference
		fmt.Printf("%#v\n", lcoCourse) // to print interfaces %#v
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where we want to add data in to key value pair

	var myOnlineData map[string]interface{}        // as I don't know the `value`` is all time string
	json.Unmarshal(jsonDataFromWeb, &myOnlineData) //we are passing reference, not copy
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
