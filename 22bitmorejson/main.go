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
	EncodeJson()
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
