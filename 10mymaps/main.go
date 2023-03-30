package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	//adding
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS short for: ", languages["JS"])

	//deleteing
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops through the maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
