package main

import "fmt"

func main() {
	number_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range number_slice {
		if number%2 == 0 {

			fmt.Println(number, " is even number")
		} else {
			fmt.Println(number, " is odd number")

		}
	}
}
