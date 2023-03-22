package main

import (
	"fmt"
)

func main() {
	//way 1
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#f0001",
	}

	//way 2
	// var colours map[string]string

	//Way 3
	// colours := make(map[string]string)

	colours["whilte"] = "#f0ff1"

	//delete a key
	// delete(colours, "whilte")
	colours["red"] = "#ff0000"

	printMap(colours)
}
func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for ", colour, "is ", hex)
	}
}
