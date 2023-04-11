package main

import (
	"fmt"
	"mongoGo/router"
)

func main() {
	fmt.Println("hello")

	router := router.Router()
	router.Run(":8080")
}
