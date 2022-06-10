package main

import "fmt"

func main() {
	var rows = 5
	for i := 0; i <= rows; i++ {
		for j := 0; j <= 5; j++ {
			if j > i {
				fmt.Printf("*")
			}

		}
		fmt.Println()
	}
}
