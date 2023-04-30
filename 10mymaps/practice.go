package main

import "fmt"

func main() {
	fmt.Println("hello")

	var abc map[string]string
	var letter []map[string]string
	var def map[string]string

	abc = make(map[string]string)
	def = make(map[string]string)
	letter = make([]map[string]string, 0)

	abc["safat"] = "safat is a good boy"
	abc["sayem"] = "sayem is a dorbesh"
	fmt.Println(abc)

	letter = append(letter, abc)
	fmt.Println(letter[0])

	def["bangladesh"] = "best country"
	def["india"] = "worst country"

	fmt.Println(letter)

	letter = append(letter, def)
	fmt.Println(letter)

}
